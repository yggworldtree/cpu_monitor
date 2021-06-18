package cmd

import (
	hbtp "github.com/mgr9525/HyperByte-Transfer-Protocol"
	"github.com/yggworldtree/cpu_monitor/monitor"
	"github.com/yggworldtree/go-sdk/ywtree"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

const Version = "0.1.1"

var (
	app       = kingpin.New("cpuMon", "A ywtree cpu monitor application.")
	HbtpHost  = ""
	DevSecret = ""
	DevName   = ""
	Debug     = false
)

func Run() {
	regs()
	kingpin.Version(Version)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
func regs() {
	app.Flag("host", "ywtree host").Short('h').Default("localhost:7000").StringVar(&HbtpHost)
	app.Flag("secret", "ywtree dev secret").Short('s').StringVar(&DevSecret)
	app.Flag("name", "ywtree dev name").Short('n').StringVar(&DevName)
	cmd := app.Command("run", "run process").Default().
		Action(func(pc *kingpin.ParseContext) error {
			return run()
		})
	cmd.Flag("debug", "debug log show").BoolVar(&Debug)
}
func run() error {
	if Debug {
		hbtp.Debug = true
	}
	mgr := monitor.NewManager(DevName)
	monitor.YwtEgn = ywtree.NewEngine(mgr, &ywtree.Config{
		Host:   HbtpHost,
		Org:    "mgr",
		Name:   "cpu-monitor",
		Secret: DevSecret,
	})
	monitor.YwtEgn.RegGrpcGrpcFun(1, &monitor.DevRPC{})
	return monitor.YwtEgn.Run()
}
