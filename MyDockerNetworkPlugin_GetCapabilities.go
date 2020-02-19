package main

import (
	n "github.com/docker/go-plugins-helpers/network"
	datastore "github.com/docker/libnetwork/datastore"
)

func (self *MyDockerNetworkPlugin) GetCapabilities() (*n.CapabilitiesResponse, error) {

	response := &n.CapabilitiesResponse{
		Scope:             datastore.LocalScope,
		ConnectivityScope: datastore.GlobalScope,
	}

	return response, nil
}
