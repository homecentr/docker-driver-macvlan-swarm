package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) EndpointInfo(req *n.InfoRequest) (*n.InfoResponse, error) {
	log.Printf("Received EndpointOperInfo req:\n%+v\n", req)
	// return empty map for now - the value of the Value field is an arbitrary (possibly empty) map
	value := make(map[string]string)
	resp := &n.InfoResponse{
		Value: value,
	}
	return resp, nil
}
