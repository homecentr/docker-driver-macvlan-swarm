package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) DiscoverNew(req *n.DiscoveryNotification) error {
	log.Printf("Received DiscoverNew req:\n%+v\n", req)
	if req.DiscoveryType == 1 {
		// Node Discovery
		log.Printf("....which is of type NodeDiscovery\n")
	}
	// used for plumbing, do nothing API for now
	return nil
}
