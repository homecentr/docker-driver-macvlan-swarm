package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) ProgramExternalConnectivity(req *n.ProgramExternalConnectivityRequest) error {
	log.Printf("Received ProgramExternalConnectivity req:\n%+v\n", string(json.Marshal(req)))

	return mac.ProgramExternalConnectivity(req)
}
