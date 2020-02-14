package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) DiscoverNew(req *n.DiscoveryNotification) error {
	log.Printf("Received DiscoverNew req:\n%+v\n", string(json.Marshal(req)))

	return mac.DiscoverNew(req)
}
