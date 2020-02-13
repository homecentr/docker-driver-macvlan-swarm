package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) CreateEndpoint(req *n.CreateEndpointRequest) (*n.CreateEndpointResponse, error) {
	log.Printf("Received CreateEndpoint req:\n%+v\n", req)

	// If the remote process was supplied a non-empty value in Interface, it
	// must respond with an empty Interface value. LibNetwork will treat it as
	// an error if it supplies a non-empty value and receives a non-empty value
	// back, and roll back the operation.
	intfInfo := new(n.EndpointInterface)

	if req.Interface == nil {
		// case never hit in docker v1.11.0, but in tests
		intfInfo.Address = "1.1.1.1/24"
		// AddressIPv6 - optional
		intfInfo.MacAddress = "00:00:00:00:00:aa"
	}
	resp := &n.CreateEndpointResponse{
		Interface: intfInfo,
	}
	return resp, nil
}
