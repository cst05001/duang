package dockerclienteng1

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/version"
)

func (this *DockerClient) Version() *types.Version {
	v1, err := this.Client.Version()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	v2 := &types.Version{
		Version:       v1.Get("Version"),
		ApiVersion:    version.Version(v1.Get("ApiVersion")),
		GitCommit:     v1.Get("GitCommit"),
		GoVersion:     v1.Get("GoVersion"),
		Os:            v1.Get("Os"),
		Arch:          v1.Get("Arch"),
		KernelVersion: v1.Get("KernelVersion"),
		Experimental:  v1.GetBool("Experimental"),
	}
	return v2
}
