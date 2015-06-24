package dockerdscheduler

import (
	"github.com/cst05001/duang/models/core"
)

type DockerdSchedulerInterface interface {
	GetDockerd(n int64) []*core.Dockerd
}
