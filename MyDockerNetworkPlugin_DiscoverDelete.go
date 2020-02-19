package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) DiscoverDelete(req *n.DiscoveryNotification) error {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received DiscoverDelete req:\n%+v\n", string(reqJson))

	return nil
}
