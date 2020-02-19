package mdnp

import (
	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) Leave(req *n.LeaveRequest) error {
	printJson("Leave", req)

	return nil
}
