package comm

import (
	"github.com/yggworldtree/go-core/bean"
)

var (
	MsgPthCpuMem = bean.NewTopicPath("system", "cpu-mem")
)

type MsgCpuInfo struct {
	Percents   []float64 `json:"percents"`
	Average    float64   `json:"average"`
	ProcessLen int       `json:"processLen"`
}
type MsgMemInfo struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
	Sin         uint64  `json:"sin"`
	Sout        uint64  `json:"sout"`
	PgIn        uint64  `json:"pgIn"`
	PgOut       uint64  `json:"pgOut"`
	PgFault     uint64  `json:"pgFault"`

	// Linux specific numbers
	// https://www.kernel.org/doc/Documentation/cgroup-v2.txt
	PgMajFault uint64 `json:"pgMajFault"`
}

type MsgBox struct {
	Name       string     `json:"name"`
	Cpu        MsgCpuInfo `json:"cpu"`
	SwapMem    MsgMemInfo `json:"swapMem"`
	VirtualMem MsgMemInfo `json:"virtualMem"`
}

type MsgCpuInfos struct {
	CPU        int32    `json:"cpu"`
	VendorID   string   `json:"vendorId"`
	Family     string   `json:"family"`
	Model      string   `json:"model"`
	Stepping   int32    `json:"stepping"`
	PhysicalID string   `json:"physicalId"`
	CoreID     string   `json:"coreId"`
	Cores      int32    `json:"cores"`
	ModelName  string   `json:"modelName"`
	Mhz        float64  `json:"mhz"`
	CacheSize  int32    `json:"cacheSize"`
	Flags      []string `json:"flags"`
	Microcode  string   `json:"microcode"`
}
type ProcInfo struct {
	Pid         int      `json:"pid"`
	User        string   `json:"user"`
	Group       string   `json:"group"`
	CommandName string   `json:"commandName"`
	CommandPath string   `json:"commandPath"`
	CommandLine []string `json:"commandLine"`
	CreateTime  int64    `json:"createTime"`
	Terminal    string   `json:"terminal"`
	Status      []string `json:"status"`
	Cwd         string   `json:"cwd"`
	Cpu         float64  `json:"cpu"`
	Mem         float64  `json:"mem"`
}
