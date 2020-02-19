package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	"github.com/docker/libnetwork/types"
)

func (self *MyDockerNetworkPlugin) FreeNetwork(req *n.FreeNetworkRequest) error {
	reqJson, _ := json.Marshal(req)
	log.Printf("Respose FreeNetwork req:\n%+v\n", string(reqJson))

	// As per the built-in implementation
	return types.NotImplementedErrorf("not implemented")
}
