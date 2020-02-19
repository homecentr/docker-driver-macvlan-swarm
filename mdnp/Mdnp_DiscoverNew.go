package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) DiscoverNew(req *n.DiscoveryNotification) error {
	printJson("DiscoverNew", req)

	return nil
}
