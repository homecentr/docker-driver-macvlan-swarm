package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) CreateEndpoint(req *n.CreateEndpointRequest) (*n.CreateEndpointResponse, error) {
	printJson("CreateEndpoint", req)

	return nil, nil
}
