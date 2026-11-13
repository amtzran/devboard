// Package server provides the HTTP server setup for the devboard API.
package server

import (
	"net/http"
	"time"
)

// Server wraps the HTTP server and its router.
type Server struct {
	httpServer *http.Server
	mux        *http.ServeMux
}

// Middleware wraps an http.Handler with additional behavior.
type Middleware func(http.Handler) http.Handler

// New creates a Server listening on addr.
func New(addr string) *Server {
	mux := http.NewServeMux()
	httpServer := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
		mux:        mux,
	}
}

// Start begins serving HTTP requests.
func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

// RegisterRoute registers handler for requests matching pattern.
func (s *Server) RegisterRoute(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

// ServeHTTP implements http.Handler.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

// Use wraps the server's handler with m. Middlewares run in reverse
// order of registration the last one added handles the request first.
// It must be called before Start.
func (s *Server) Use(m Middleware) {
	s.httpServer.Handler = m(s.httpServer.Handler)
}
