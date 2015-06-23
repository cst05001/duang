package models

import (
	"github.com/cst05001/duang/models/core"
	scheduler1 "github.com/cst05001/duang/models/dockerdscheduler/scheduler1"
)

var Scheduler core.DockerdSchedulerInterface

func init() {

	Scheduler = scheduler1.NewDockerdScheduler1()
}