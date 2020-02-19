package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) ProgramExternalConnectivity(req *n.ProgramExternalConnectivityRequest) error {
	printJson("ProgramExternalConnectivity", req)

	return nil
}
