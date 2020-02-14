package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) FreeNetwork(req *n.FreeNetworkRequest) error {
	log.Printf("Respose FreeNetwork req:\n%+v\n", string(json.Marshal(response)))

	return mac.FreeNetwork(req)
}
