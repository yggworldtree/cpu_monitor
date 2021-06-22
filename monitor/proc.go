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
		exe, _ := v.Exe()
		cmdlines, _ := v.CmdlineSlice()
		cwd, _ := v.Cwd()
		stats, _ := v.Status()
		cpup, _ := v.CPUPercent()
		memp, _ := v.MemoryPercent()
		user, _ := v.Username()
		term, _ := v.Terminal()
		ctm, _ := v.CreateTime()

		t := &comm.ProcInfo{
			Pid:         int(pid),
			User:        user,
			CommandName: command,
			CommandPath: exe,
			CommandLine: cmdlines,
			CreateTime:  ctm,
			Terminal:    term,
			Status:      stats,
			Cwd:         cwd,
			Cpu:         cpup,
			Mem:         float64(memp),
		}
		ls = append(ls, t)
	}

	return ls, nil
}
