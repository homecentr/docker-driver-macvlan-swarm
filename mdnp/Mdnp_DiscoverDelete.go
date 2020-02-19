package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) DiscoverDelete(req *n.DiscoveryNotification) error {
	printJson("DiscoverDelete", req)

	return nil
}
