package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) CreateNetwork(req *n.CreateNetworkRequest) error {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received CreateNetwork req:\n%+v\n", string(reqJson))

	return nil
}
