package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

func Respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error encoding response as JSON"))
		log.Printf("error encoding response as JSON: %v", err)
	}
}

type ErrorMessage struct {
	Message string
}

func Error(w http.ResponseWriter, r *http.Request, status int, msg string) {
	ve := ErrorMessage{Message: msg}
	Respond(w, r, status, ve)
}
