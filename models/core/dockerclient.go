package core

import ()

type DockerClient interface {
	Run(*Unit) error
}
