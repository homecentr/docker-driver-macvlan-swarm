package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) Join(req *n.JoinRequest) (*n.JoinResponse, error) {
	printJson("Join", req)

	return nil, nil
}
