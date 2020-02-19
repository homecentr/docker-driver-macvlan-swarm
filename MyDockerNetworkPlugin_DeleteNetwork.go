package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) DeleteNetwork(req *n.DeleteNetworkRequest) error {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received DeleteNetwork req:\n%+v\n", string(reqJson))

	return nil
}
