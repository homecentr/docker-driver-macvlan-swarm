package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) RevokeExternalConnectivity(req *n.RevokeExternalConnectivityRequest) error {
	log.Printf("Received RevokeExternalConnectivity req:\n%+v\n", string(json.Marshal(req)))

	return mac.RevokeExternalConnectivity(req)
}
