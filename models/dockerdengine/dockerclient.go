package dockerdengine

import (
	"github.com/cst05001/duang/models/core"
)

type DockerClient interface {
	//interface1 传入回调函数
	Run(*core.Unit, func(*core.Dockerd)) error
}
