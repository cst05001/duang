package models

import (
	"github.com/cst05001/duang/models/dockerdscheduler"
	scheduler1 "github.com/cst05001/duang/models/dockerdscheduler/scheduler1"
)

var Scheduler dockerdscheduler.DockerdSchedulerInterface

func init() {

	Scheduler = scheduler1.NewDockerdScheduler1()
}