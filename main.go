// main package contains the implementation of my docker network plugin, a.k.a.
// mdnp
package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	m "github.com/homecentr/docker-driver-macvlan-swarm/mdnp"

	n "github.com/docker/go-plugins-helpers/network"
)

var (
	PLUGIN_NAME = "mdnp"
)

// return the default prod config for ad_service
func NewMyDockerNetworkPlugin(scope string) (*m.MyDockerNetworkPlugin, error) {
	mdnp := &m.MyDockerNetworkPlugin{
		Scope: scope,
	}
	return mdnp, nil
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	driver, err := NewMyDockerNetworkPlugin("local")
	if err != nil {
		log.Fatalf("ERROR: %s init failed!", PLUGIN_NAME)
	}

	requestHandler := n.NewHandler(driver)

	// err = os.RemoveAll(s.WindowsDefaultDaemonRootDir() + "\\plugins")
	// if err != nil {
	// 	log.Fatalf("Removing dir failed...")
	// }

	err1 := requestHandler.ServeUnix(PLUGIN_NAME, 0)
	if err1 != nil {
		log.Fatalf("FATAL ERROR.... %s", err1)
	}
}
