package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) CreateEndpoint(req *n.CreateEndpointRequest) (*n.CreateEndpointResponse, error) {
	json, _ := json.Marshal(req)
	log.Printf("Received CreateEndpoint req:\n%+v\n", string(json))

	response, err = mac.CreateEndpoint(req)

	json, _ = json.Marshal(response)
	log.Printf("Returning CreateEndpoint respone:\n%+v\n", string(json))

	return response, err
}
