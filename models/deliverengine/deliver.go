package deliverengine

import (
	"github.com/cst05001/duang/models/core"
)

type DeliverInterface interface {
	AddBackend(name string, backends []string) error
	DelBackend(name string, backends []string) error
	AddFrontend(*core.Frontend) (*core.Frontend, error)
	DelFrontend(*core.Frontend) error
}
