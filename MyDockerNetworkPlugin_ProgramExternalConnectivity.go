package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) ProgramExternalConnectivity(req *n.ProgramExternalConnectivityRequest) error {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received ProgramExternalConnectivity req:\n%+v\n", string(reqJson))

	return nil
}
