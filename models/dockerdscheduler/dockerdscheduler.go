package dockerdscheduler

import (
	"github.com/cst05001/duang/models/core"
)

type DockerdSchedulerInterface interface {
	//n, excludeBackends
	GetDockerd(int64, []string) []*core.Dockerd
}
