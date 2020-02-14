package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) DeleteEndpoint(req *n.DeleteEndpointRequest) error {
	log.Printf("Received DeleteEndpoint req:\n%+v\n", string(json.Marshal(req)))

	err = mac.CreateEndpoint(req)

	return err
}
