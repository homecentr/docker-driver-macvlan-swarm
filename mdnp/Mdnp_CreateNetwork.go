package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) CreateNetwork(req *n.CreateNetworkRequest) error {
	printJson("CreateNetwork", req)

	return nil
}
