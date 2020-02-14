package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) AllocateNetwork(req *n.AllocateNetworkRequest) (*n.AllocateNetworkResponse, error) {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received AllocateNetwork req:\n%+v\n", string(reqJson))

	response, err := mac.AllocateNetwork(req)

	respJson, _ = json.Marshal(response)
	log.Printf("Returning AllocateNetwork respone:\n%+v\n", string(respJson))

	return response, err
}
