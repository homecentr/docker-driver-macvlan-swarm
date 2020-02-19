package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) FreeNetwork(req *n.FreeNetworkRequest) error {
	printJson("FreeNetwork", req)

	return nil
}
