package main

import (
	"encoding/json"
	"log"

	n "github.com/docker/go-plugins-helpers/network"
	mac "github.com/docker/libnetwork/drivers/macvlan"
)

func (self *MyDockerNetworkPlugin) EndpointInfo(req *n.InfoRequest) (*n.InfoResponse, error) {
	log.Printf("Received EndpointOperInfo req:\n%+v\n", string(json.Marshal(req)))

	response, err = mac.EndpointInfo(req)

	log.Printf("Respose EndpointOperInfo req:\n%+v\n", string(json.Marshal(response)))

	return response, err
}
