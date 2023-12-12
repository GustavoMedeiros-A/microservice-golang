package main

import (
	"net/http"
)

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "hit the broker",
	}
	// Just this line of code, now do the same of the comment code
	_ = app.writeJSON(w, http.StatusOK, payload)

	// _ -> means: "ignore the error"
	// Marshal - converting Go objects (such as structs) into a serialized format like JSON, XML, YAML
	// out, _ := json.MarshalIndent(payload, "", "\t")
	// w.Header().Set("Content-Type", "application/json") // Set Content-Type to application/json
	// w.WriteHeader(http.StatusAccepted)
	// w.Write(out)

}
