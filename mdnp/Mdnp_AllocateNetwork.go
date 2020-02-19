package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
	"github.com/docker/libnetwork/types"
)

func (self *MyDockerNetworkPlugin) AllocateNetwork(req *n.AllocateNetworkRequest) (*n.AllocateNetworkResponse, error) {
	printJson("AllocateNetwork", req)

	// As per the built-in implementation
	return nil, types.NotImplementedErrorf("not implemented")
}
