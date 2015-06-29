package deliverengine

import (
	"github.com/cst05001/duang/models/core"
)

type DeliverInterface interface {
	AddBackend(name string, backend *core.Backend) (*core.Backend, error)
	DelBackend(name string, backend *core.Backend) error
	AddFrontend(*core.Frontend) (*core.Frontend, error)
	DelFrontend(*core.Frontend) error
}
