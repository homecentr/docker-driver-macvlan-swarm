package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) GetCapabilities() (*n.CapabilitiesResponse, error) {

	response, err = mac.GetCapabilities()

	log.Printf("Respose GetCapabilities req:\n%+v\n", string(json.Marshal(response)))

	return response, error
}
