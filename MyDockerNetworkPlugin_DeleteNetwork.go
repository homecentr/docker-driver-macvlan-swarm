package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) DeleteNetwork(req *n.DeleteNetworkRequest) error {
	log.Printf("Received DeleteNetwork req:\n%+v\n", string(json.Marshal(req)))

	return mac.DeleteNetwork(req)
}
