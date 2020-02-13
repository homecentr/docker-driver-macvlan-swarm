package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) CreateNetwork(req *n.CreateNetworkRequest) error {
	log.Printf("Received CreateNetwork req:\n%+v\n", req)
	// used for plumbing, do nothing API for now
	return nil
}
