module github.com/yggworldtree/cpu_monitor

go 1.16

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20210208195552-ff826a37aa15 // indirect
	github.com/mgr9525/HyperByte-Transfer-Protocol v1.1.5
	github.com/shirou/gopsutil/v3 v3.21.5
	github.com/yggworldtree/go-core v0.0.0-20210621070134-be038d4e6f72
	github.com/yggworldtree/go-sdk v0.0.0-20210621073455-b0231047e196
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
)

replace (
	github.com/yggworldtree/go-core => ../go-core
	github.com/yggworldtree/go-sdk => ../go-sdk
)
