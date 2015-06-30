package deliverengine

import ()

type DeliverInterface interface {
	Bind(frontend, domain string, backend string) error
}
