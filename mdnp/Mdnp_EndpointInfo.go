package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) EndpointInfo(req *n.InfoRequest) (*n.InfoResponse, error) {
	printJson("EndpointInfo", req)

	return nil, nil
}
