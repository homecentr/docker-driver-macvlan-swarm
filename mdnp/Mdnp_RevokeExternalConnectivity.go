package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) RevokeExternalConnectivity(req *n.RevokeExternalConnectivityRequest) error {
	printJson("RevokeExternalConnectivity", req)

	return nil
}
