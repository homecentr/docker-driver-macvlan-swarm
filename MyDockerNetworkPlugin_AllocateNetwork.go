package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) AllocateNetwork(req *n.AllocateNetworkRequest) (*n.AllocateNetworkResponse, error) {
	log.Printf("Received Allocate network req:\n%+v\n", req)
	// used for plumbing, do nothing API for now
	return nil, nil
}
