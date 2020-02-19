package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	"github.com/docker/libnetwork/types"
)

func (self *MyDockerNetworkPlugin) AllocateNetwork(req *n.AllocateNetworkRequest) (*n.AllocateNetworkResponse, error) {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received AllocateNetwork req:\n%+v\n", string(reqJson))

	// As per the built-in implementation
	return nil, types.NotImplementedErrorf("not implemented")
}
