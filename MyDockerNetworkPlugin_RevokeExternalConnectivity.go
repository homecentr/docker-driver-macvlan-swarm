package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) RevokeExternalConnectivity(req *n.RevokeExternalConnectivityRequest) error {
	reqJson, _ := json.Marshal(req)
	log.Printf("Received RevokeExternalConnectivity req:\n%+v\n", string(reqJson))

	return nil
}
