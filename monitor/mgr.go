package monitor

import (
	"bytes"
	"context"
	"fmt"
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/yggworldtree/cpu_monitor/comm"
	"github.com/yggworldtree/go-core/utils"
	"github.com/yggworldtree/go-sdk/ywtree"
	"sync"
	"time"
)

var YwtEgn *ywtree.Engine

type Manager struct {
	name string
	Ctx  context.Context
	cncl context.CancelFunc

	msgtmr *utils.Timer

	cpulk   sync.Mutex
	cpuinfo comm.MsgCpuInfo
}

func NewManager(name string) *Manager {
	c := &Manager{
		name:   name,
		msgtmr: utils.NewTimer(time.Second + time.Millisecond*50),
	}
	c.Ctx, c.cncl = context.WithCancel(context.Background())
	go func() {
		for !hbtp.EndContext(c.Ctx) {
			c.runCpu()
			time.Sleep(time.Millisecond)
		}
	}()
	go func() {
		for !hbtp.EndContext(c.Ctx) {
			c.runMsg()
			time.Sleep(time.Millisecond)
		}
		if YwtEgn != nil {
			YwtEgn.Stop()
		}
	}()
	return c
}
func (c *Manager) runCpu() {
	defer func() {
		if err := recover(); err != nil {
			hbtp.Debugf("Manager runCpu recover:%v", err)
		}
	}()

	cs, err := cpu.PercentWithContext(c.Ctx, time.Second, true)
	if err != nil {
		hbtp.Debugf("cpu.Percent err:%v", err)
		return
	}
	ln := len(cs)
	if ln <= 0 {
		hbtp.Debugf("cpu.Percent cpu len is 0?")
		return
	}

	c.cpulk.Lock()
	defer c.cpulk.Unlock()
	c.cpuinfo.Percents = cs
	c.cpuinfo.Average = 0
	outs := &bytes.Buffer{}
	for i, v := range cs {
		outs.WriteString(fmt.Sprintf("[%d]:%.4f%%, ", i+1, v))
		c.cpuinfo.Average += v
	}
	hbtp.Debugf("cpu.Percent:" + outs.String())
	c.cpuinfo.Average = c.cpuinfo.Average / float64(ln)
}
func (c *Manager) runMsg() {
	defer func() {
		if err := recover(); err != nil {
			hbtp.Debugf("Manager runMsg recover:%v", err)
		}
	}()
	if !c.msgtmr.Tick() {
		return
	}
	v1, err := mem.SwapMemoryWithContext(c.Ctx)
	if err != nil {
		hbtp.Debugf("mem.SwapMemory err:%v", err)
		return
	}
	v2, err := mem.VirtualMemoryWithContext(c.Ctx)
	if err != nil {
		hbtp.Debugf("mem.SwapMemory err:%v", err)
		return
	}
	hbtp.Debugf("mem.totalMemory Swap:%.4f%%, Virtual:%.4f%%", v1.UsedPercent, v2.UsedPercent)
	box := &comm.MsgBox{
		Name: c.name,
		Cpu:  c.cpuinfo,
		SwapMem: comm.MsgMemInfo{
			Total:       v1.Total,
			Used:        v1.Used,
			Free:        v1.Free,
			UsedPercent: v1.UsedPercent,
			Sin:         v1.Sin,
			Sout:        v1.Sout,
			PgIn:        v1.PgIn,
			PgOut:       v1.PgOut,
			PgFault:     v1.PgFault,
			PgMajFault:  v1.PgMajFault,
		},
		VirtualMem: comm.MsgMemInfo{
			Total:       v2.Total,
			Used:        v2.Used,
			Free:        v2.Free,
			UsedPercent: v2.UsedPercent,
		},
	}
	err = YwtEgn.PushTopic(comm.MsgPthCpuMem, box)
	if err != nil {
		hbtp.Debugf("PushTopic err:%v", err)
	}
}
