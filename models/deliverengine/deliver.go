package deliverengine

type DeliverInterface interface {
	AddBackend(frontend string, backends []string) error
	DelBackend(frontend string, backends []string) error
	AddFrontend(frontend string) error
	DelFrontend(frontend string) error
}
