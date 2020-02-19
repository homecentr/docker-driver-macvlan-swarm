package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) Join(req *n.JoinRequest) (*n.JoinResponse, error) {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received Join req:\n%+v\n", string(reqJson))

	// TODO: Connect container to the impl network

	return nil, nil
}
