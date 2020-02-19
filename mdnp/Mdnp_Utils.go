package mdnp

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func printJson(caller string, obj interface{}) {
	objJson, err := json.Marshal(obj)
	if err != nil {
		log.WithError(err).Error("Conversion to JSON failed (" + caller + ") ")
	}

	log.WithField("Request/Response", string(objJson)).Info("Call handled in " + caller)
}
