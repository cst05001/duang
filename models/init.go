package models

import (
	"github.com/cst05001/duang/models/core"
	"github.com/cst05001/duang/models/dockerdscheduler"
	mcscheduler "github.com/cst05001/duang/models/dockerdscheduler/MCscheduler"
)

var Scheduler dockerdscheduler.DockerdSchedulerInterface
var IPPool *core.IpPool

func init() {

	Scheduler = mcscheduler.NewScheduler()
	IPPool = core.NewIpPool()
}
