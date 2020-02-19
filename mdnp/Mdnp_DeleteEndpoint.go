package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) DeleteEndpoint(req *n.DeleteEndpointRequest) error {
	printJson("DeleteEndpoint", req)

	return nil
}
