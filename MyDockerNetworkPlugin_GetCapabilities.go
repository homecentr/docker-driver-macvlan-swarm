package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) GetCapabilities() (*n.CapabilitiesResponse,
	error) {
	log.Printf("Received GetCapabilities req")
	capabilities := &n.CapabilitiesResponse{
		Scope: self.scope,
	}
	return capabilities, nil
}
