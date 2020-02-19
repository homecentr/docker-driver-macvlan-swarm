package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) DeleteEndpoint(req *n.DeleteEndpointRequest) error {
	reqJson, _ := json.Marshal(req)
	log.Printf("DeleteEndpoint.Request:\n%+v\n", string(reqJson))

	return nil
}
