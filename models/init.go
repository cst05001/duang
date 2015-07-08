package models

import (
	"github.com/cst05001/duang/models/core"
	"github.com/cst05001/duang/models/dockerdengine"
	dockerd_engine1 "github.com/cst05001/duang/models/dockerdengine/engine1"
	"github.com/cst05001/duang/models/dockerdscheduler"
	mcscheduler "github.com/cst05001/duang/models/dockerdscheduler/MCscheduler"
)

var Scheduler dockerdscheduler.DockerdSchedulerInterface
var IPPool *core.IpPool
var DockerClient dockerdengine.DockerClient

func init() {
	DockerClient = dockerd_engine1.NewDockerClient()
	Scheduler = mcscheduler.NewScheduler()
	IPPool = core.NewIpPool()
}
