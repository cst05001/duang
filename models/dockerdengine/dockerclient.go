package dockerdengine

import (
	"github.com/cst05001/duang/models/core"
)

const (
	STATUS_ON_CREATE_SUCCESSED = iota
	STATUS_ON_CREATE_FAILED
	STATUS_ON_RUN_SUCCESSED
	STATUS_ON_RUN_FAILED
)

const (
	STATUS_CONTAINER_DOWN = iota
	STATUS_CONTAINER_UP
	STATUS_CONTAINER_NOEXIST
)

type DockerClient interface {
	//Unit, calllbackFunc(*core.Dockerd, int type status, ...args)
	Run(*core.Unit, func(*core.Dockerd, int, ...interface{})) error
	UpdateContainerStatus(*core.Unit) map[*core.Dockerd]uint8
}
