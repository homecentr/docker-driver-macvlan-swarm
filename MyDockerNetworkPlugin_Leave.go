package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) Leave(req *n.LeaveRequest) error {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received Leave req:\n%+v\n", string(reqJson))

	return nil
}
