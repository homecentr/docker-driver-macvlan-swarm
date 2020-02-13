package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) RevokeExternalConnectivity(req *n.RevokeExternalConnectivityRequest) error {
	log.Printf("Received RevokeExternalConnectivity req:\n%+v\n", req)
	// no documentation?!
	// used for plumbing, do nothing API for now
	return nil
}
