package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) Join(req *n.JoinRequest) (*n.JoinResponse, error) {
	log.Printf("Received Join req:\n%+v\n", string(json.Marshal(req)))

	response, err = mac.Join(req)

	log.Printf("Respose Join req:\n%+v\n", string(json.Marshal(response)))

	return response, error
}
