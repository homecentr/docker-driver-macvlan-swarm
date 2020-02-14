package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) AllocateNetwork(req *n.AllocateNetworkRequest) (*n.AllocateNetworkResponse, error) {
	json, _ := json.Marshal(req)
	log.Printf("Received AllocateNetwork req:\n%+v\n", string(json))

	response, err := mac.AllocateNetwork(req)

	json, _ = json.Marshal(response)
	log.Printf("Returning AllocateNetwork respone:\n%+v\n", string(json))

	return response, err
}
