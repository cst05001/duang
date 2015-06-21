package models

import (
	"github.com/docker/docker/api/types"
)

type dockerclient interface {
	Version() types.Version
	Info() types.Info
}
