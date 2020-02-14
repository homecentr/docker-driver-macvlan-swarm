package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) Leave(req *n.LeaveRequest) error {
	log.Printf("Received Leave req:\n%+v\n", string(json.Marshal(req)))

	return mac.Leave(req)
}
