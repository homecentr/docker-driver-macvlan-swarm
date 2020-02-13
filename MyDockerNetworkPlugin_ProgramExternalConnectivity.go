package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) ProgramExternalConnectivity(req *n.ProgramExternalConnectivityRequest) error {
	log.Printf("Received ProgramExternalConnectivity req:\n%+v\n", req)
	// no documentation?!
	// used for plumbing, do nothing API for now
	return nil
}
