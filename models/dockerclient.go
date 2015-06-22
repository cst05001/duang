package models

import ()

type DockerClient interface {
	Run(*Unit) error
}
