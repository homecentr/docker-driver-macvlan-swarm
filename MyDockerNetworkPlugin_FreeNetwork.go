package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) FreeNetwork(req *n.FreeNetworkRequest) error {
	log.Printf("Received Free network req:\n%+v\n", req)
	// used for plumbing, do nothing API for now
	return nil
}
