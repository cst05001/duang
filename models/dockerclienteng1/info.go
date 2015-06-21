package dockerclienteng1

import (
	"fmt"
	"github.com/docker/docker/api/types"
)

func (this *DockerClient) Info() *types.Info {
	i1, err := this.Client.Info()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	i2 := &types.Info{
		ID:         i1.Get("ID"),
		Containers: i1.GetInt("Containers"),
		Images:     i1.GetInt("Images"),
		Driver:     i1.Get("Driver"),
		//DriverStatus:      [][2]string
		MemoryLimit:        i1.GetBool("MemoryLimit"),
		SwapLimit:          i1.GetBool("SwapLimit"),
		CpuCfsPeriod:       i1.GetBool("CpuCfsPeriod"),
		CpuCfsQuota:        i1.GetBool("CpuCfsQuota"),
		IPv4Forwarding:     i1.GetBool("IPv4Forwarding"),
		BridgeNfIptables:   i1.GetBool("BridgeNfIptables"),
		BridgeNfIp6tables:  i1.GetBool("BridgeNfIp6tables"),
		Debug:              i1.GetBool("Debug"),
		NFd:                i1.GetInt("NFd"),
		OomKillDisable:     i1.GetBool("OomKillDisable"),
		NGoroutines:        i1.GetInt("NGoroutines"),
		SystemTime:         i1.Get("SystemTime"),
		ExecutionDriver:    i1.Get("ExecutionDriver"),
		LoggingDriver:      i1.Get("LoggingDriver"),
		NEventsListener:    i1.GetInt("NEventsListener"),
		KernelVersion:      i1.Get("KernelVersion"),
		OperatingSystem:    i1.Get("OperatingSystem"),
		IndexServerAddress: i1.Get("IndexServerAddress"),
		//RegistryConfig:    interface{}
		InitSha1:      i1.Get("InitSha1"),
		InitPath:      i1.Get("InitPath"),
		NCPU:          i1.GetInt("NCPU"),
		MemTotal:      i1.GetInt64("MemTotal"),
		DockerRootDir: i1.Get("DockerRootDir"),
		HttpProxy:     i1.Get("HttpProxy"),
		HttpsProxy:    i1.Get("HttpsProxy"),
		NoProxy:       i1.Get("NoProxy"),
		Name:          i1.Get("Name"),
		//Labels:            []string
		ExperimentalBuild: i1.GetBool("ExperimentalBuild"),
	}
	fmt.Println(i2)
	return i2
}
