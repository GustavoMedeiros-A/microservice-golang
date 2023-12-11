package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "hit the broker",
	}
	// _ -> means: "ignore the error"
	// Marshal - converting Go objects (such as structs) into a serialized format like JSON, XML, YAML
	out, _ := json.MarshalIndent(payload, "", "\t")
	w.Header().Set("Content-Type", "application/json") // Set Content-Type to application/json
	w.WriteHeader(http.StatusAccepted)
	w.Write(out)

}
