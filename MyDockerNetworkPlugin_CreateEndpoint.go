package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) CreateEndpoint(req *n.CreateEndpointRequest) (*n.CreateEndpointResponse, error) {
	reqJson, _ := json.Marshal(req)
	log.Printf("AllocateNetwork.Request:\n%+v\n", string(reqJson))

	return nil, nil
}
