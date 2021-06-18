package monitor

import (
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/yggworldtree/cpu_monitor/comm"
	"time"
)

type DevRPC struct{}

func (cs *DevRPC) AuthFun() hbtp.AuthFun {
	return nil
}
func (cs *DevRPC) CpuInfos(c *hbtp.Context) {
	tms := time.Now()
	defer func() {
		hbtp.Debugf("DevRPC CpuInfos times:%.4fs", time.Since(tms).Seconds())
	}()
	var infos []*comm.MsgCpuInfos
	info, err := cpu.Info()
	if err != nil {
		c.ResString(hbtp.ResStatusErr, "info err:"+err.Error())
		return
	}
	for _, v := range info {
		infos = append(infos, &comm.MsgCpuInfos{
			CPU:        v.CPU,
			VendorID:   v.VendorID,
			Family:     v.Family,
			Model:      v.Model,
			Stepping:   v.Stepping,
			PhysicalID: v.PhysicalID,
			CoreID:     v.CoreID,
			Cores:      v.Cores,
			ModelName:  v.ModelName,
			Mhz:        v.Mhz,
			CacheSize:  v.CacheSize,
			Flags:      v.Flags,
			Microcode:  v.Microcode,
		})
	}
	c.ResJson(hbtp.ResStatusOk, infos)
}

func (cs *DevRPC) Process(c *hbtp.Context) {
	ls, err := getProcs()
	if err != nil {
		c.ResString(hbtp.ResStatusErr, "get err:"+err.Error())
		return
	}

	c.ResJson(hbtp.ResStatusOk, ls)
}
