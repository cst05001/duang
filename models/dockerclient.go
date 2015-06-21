package models

import (
	"github.com/docker/docker/api/types"
)

type dockerclient interface {
	Version() types.Version
	Info() types.Info
	CreateContainer(*Unit) (types.ContainerCreateResponse, error)
	StartContainer(string, *Unit) error
}
