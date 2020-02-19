package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) DiscoverNew(req *n.DiscoveryNotification) error {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received DiscoverNew req:\n%+v\n", string(reqJson))

	return nil
}
