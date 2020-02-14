package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) CreateNetwork(req *n.CreateNetworkRequest) error {
	json, _ := json.Marshal(req)
	log.Printf("Received CreateNetwork req:\n%+v\n", string(json))

	return mac.CreateNetwork(req)
}
