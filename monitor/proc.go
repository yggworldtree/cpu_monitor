package monitor

import (
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/yggworldtree/cpu_monitor/comm"
)

func getProcs() ([]*comm.ProcInfo, error) {
	defer func() {
		if err := recover(); err != nil {
			hbtp.Debugf("getProcs err:%v", err)
		}
	}()
	ps, err := process.Processes()
	if err != nil {
		hbtp.Debugf("failed to execute 'ps' command: %v", err)
		return nil, err
	}

	var ls []*comm.ProcInfo
	for _, v := range ps {
		pid := v.Pid
		command, _ := v.Name()
		cpup, _ := v.CPUPercent()
		memp, _ := v.MemoryPercent()
		user, _ := v.Username()

		t := &comm.ProcInfo{
			Pid:         int(pid),
			User:        user,
			CommandName: command,
			Cpu:         cpup,
			Mem:         float64(memp),
			// getting command args using gopsutil's Cmdline and CmdlineSlice wasn't
			// working the last time I tried it, so we're just reusing 'command'
			FullCommand: command,
		}
		ls = append(ls, t)
	}

	return ls, nil
}
