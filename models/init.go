package models

import (
	"github.com/cst05001/duang/models/dockerdscheduler"
	mcscheduler "github.com/cst05001/duang/models/dockerdscheduler/MCscheduler"
)

var Scheduler dockerdscheduler.DockerdSchedulerInterface

func init() {

	Scheduler = mcscheduler.NewScheduler()
}
