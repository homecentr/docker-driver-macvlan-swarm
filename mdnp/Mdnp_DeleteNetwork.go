package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) DeleteNetwork(req *n.DeleteNetworkRequest) error {
	printJson("DeleteNetwork", req)

	return nil
}
