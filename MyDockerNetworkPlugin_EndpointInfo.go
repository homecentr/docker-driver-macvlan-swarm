package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) EndpointInfo(req *n.InfoRequest) (*n.InfoResponse, error) {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received EndpointOperInfo req:\n%+v\n", string(reqJson))

	return nil, nil
}
