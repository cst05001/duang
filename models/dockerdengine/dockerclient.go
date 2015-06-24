package dockerdengine

import (
	"github.com/cst05001/duang/models/core"
)

type DockerClient interface {
	Run(*core.Unit) error
}
