package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

//Error returns bad status codes and an error in json format body
func Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	logrus.Error(err)
	Respond(w, r, code, map[string]string{"error": err.Error()})
}

//Respond returns successful status codes, response body in json
func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	logrus.Info("Status: ", code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
