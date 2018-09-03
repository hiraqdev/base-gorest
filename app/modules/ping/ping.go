package ping

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type message struct {
	Msg string `json:"message"`
}

// Handler used as ping main controller
func Handler(w http.ResponseWriter, r *http.Request) {
	output := message{"hello world"}
	j, err := json.Marshal(output)

	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("Error marshalling json")
	}

	w.Write(j)
}
