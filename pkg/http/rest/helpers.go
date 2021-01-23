package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) respond(w http.ResponseWriter, r *http.Request, data interface{}, status int) {
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

func (s *Server) error(w http.ResponseWriter, r *http.Request, status int, msg string) {
	ve := ErrorMessage{
		Message: msg,
	}
	s.respond(w, r, ve, status)
}
