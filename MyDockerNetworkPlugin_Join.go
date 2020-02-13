package main

import (
	"log"

	n "github.com/docker/go-plugins-helpers/network"
)

func (self *MyDockerNetworkPlugin) Join(req *n.JoinRequest) (*n.JoinResponse, error) {
	log.Printf("Received Join req:\n%+v\n", req)
	resp := &n.JoinResponse{
		InterfaceName: n.InterfaceName{
			// SrcName is the name of the OS level interface that the remote
			// process created
			SrcName:   "veth0",
			DstPrefix: "krish-mdnp",
		},
		// my local veth IP
		//Gateway: "172.16.0.205",
		// GatewayIPv6 - optional
		//StaticRoutes: []*n.StaticRoute{
		//	{
		//		Destination: "8.8.8.8/8",
		//		RouteType:   1,
		//		//RouteType:   0,
		//		//NextHop:     "172.16.0.205",
		//	},
		//},
		// confusion on what DisableGatewayService does - TODO google this!
		DisableGatewayService: false,
	}
	return resp, nil
}
