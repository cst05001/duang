package deliverengine

type DeliverInterface interface {
	AddBackend(name string, backends []string) error
	DelBackend(name string, backends []string) error
	AddFrontend(name, bind string) error
	DelFrontend(name, bind string) error
}
