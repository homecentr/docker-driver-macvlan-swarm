package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) Leave(req *n.LeaveRequest) error {
	log.Printf("Received Leave req:\n%+v\n", req)
	// used for plumbing, do nothing API for now
	return nil
}
