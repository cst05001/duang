package deliverengine

type DeliverInterface interface {
	SetBackend(frontend string, backends []string) error
	AddBackend(frontend string, backends []string) error
	DelBackend(frontend string, backends []string) error
	GetBackend(frontend string)		[]string
	GetFrontend()					[]string
	PauseFrontend(frontend string)	error
	ResumeFrontend(frontend string)	error
	DelFrontend(frontend string)	error
}