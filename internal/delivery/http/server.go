package http

import (
	"net/http"

	"firebase/pkg/grace"
)

// firebaseHandler ...
type firebaseHandler interface {
	// Masukkan fungsi handler di sini
	FirebaseHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server   *http.Server
	Firebase firebaseHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	return grace.Serve(port, s.Handler())
}
