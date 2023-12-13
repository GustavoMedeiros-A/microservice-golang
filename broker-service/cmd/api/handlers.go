package main

import (
	"net/http"
	"errors"
	"bytes"
	"encoding/json"
)


type RequestPayload struct {
	Action string `json:"action"`
	Auth AuthPayload `json:"auth,omitempty"`
}

type AuthPayload struct {
	Email string `json:"email"`
	Password string `json:"password"`

}

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


func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	// $ -> into "requestPayload" 
	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return 
	}

	switch requestPayload.Action {
	case "auth":
		app.authenticate(w, requestPayload.Auth)
	default:
		app.errorJSON(w, errors.New("unknown action"))
	}

}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// create some json we'll send to the auth microservice
	jsonData, _ := json.MarshalIndent(a, "", "\t")
	
	// call the auth-service to logged in
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))
	if err != nil {
		app.errorJSON(w, err)
		return 
	}
	
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		app.errorJSON(w, err)
		return 
	}
	defer response.Body.Close()

	// make sure we get back the correct status code
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("invalid credentials"))
		return 
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	// create a variable we'll read response.Body into

	var jsonFromAuthService jsonResponse

	// Decode the json
	err = json.NewDecoder(response.Body).Decode(&jsonFromAuthService)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// if error is true
	if jsonFromAuthService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return 
	}

	var payload jsonResponse
	payload.Error = false
	payload.Message = "Authenticated!"
	payload.Data = jsonFromAuthService.Data

	app.writeJSON(w, http.StatusAccepted, payload)


}