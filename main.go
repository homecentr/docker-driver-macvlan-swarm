// main package contains the implementation of my docker network plugin, a.k.a.
// mdnp
package main

import (
	"log"
	"os"

	n "github.com/docker/go-plugins-helpers/network"
	s "github.com/docker/go-plugins-helpers/sdk"
)

var (
	PLUGIN_NAME = "mdnp"
)

// return the default prod config for ad_service
func NewMyDockerNetworkPlugin(scope string) (*MyDockerNetworkPlugin, error) {
	mdnp := &MyDockerNetworkPlugin{
		scope: scope,
	}
	return mdnp, nil
}

func main() {
	driver, err := NewMyDockerNetworkPlugin("local")
	if err != nil {
		log.Fatalf("ERROR: %s init failed!", PLUGIN_NAME)
	}

	requestHandler := n.NewHandler(driver)
	log.Printf("Hello 3")

	err = os.RemoveAll(s.WindowsDefaultDaemonRootDir() + "\\plugins")
	if err != nil {
		log.Fatalf("Removing dir failed...")
	}

	err1 := requestHandler.ServeTCP(PLUGIN_NAME, ":2804", s.WindowsDefaultDaemonRootDir(), nil)
	if err1 != nil {
		log.Fatalf("FATAL ERROR.... %s", err1)
	}
}
