// Package svc is a thin HTTP server wrapper used by dwmblocks for any
// remote-control endpoint. Each feature registers its own routes via Route
// values, then the server is started with Start.
package svc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// Route is the unit of registration: one HTTP method + path + handler.
// Path uses the Go 1.22+ ServeMux pattern (without the leading method),
// e.g. "/notify" or "/items/{id}".
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// Server is the lifecycle-managed HTTP server.
type Server struct {
	addr   string
	routes []Route
	srv    *http.Server
}

func New(addr string) *Server { return &Server{addr: addr} }

func (s *Server) Handle(r ...Route) *Server {
	s.routes = append(s.routes, r...)
	return s
}

// Start binds the listener and serves in a background goroutine.
// It returns once the listener is open (or an error if it could not bind).
func (s *Server) Start() error {
	mux := http.NewServeMux()
	for _, r := range s.routes {
		mux.HandleFunc(r.Method+" "+r.Path, r.Handler)
	}
	s.srv = &http.Server{
		Addr:              s.addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("svc server: %s\n", err)
		}
	}()
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.srv == nil {
		return nil
	}
	return s.srv.Shutdown(ctx)
}

// ---------- helpers for handler authors ----------

const DefaultMaxBodyBytes = 4 << 10

// DecodeJSON enforces Content-Type: application/json, caps body size,
// and decodes into dst. On failure it writes a 400 and returns the error.
func DecodeJSON(w http.ResponseWriter, r *http.Request, dst any, maxBytes int64) error {
	if maxBytes <= 0 {
		maxBytes = DefaultMaxBodyBytes
	}
	if ct := r.Header.Get("Content-Type"); ct != "" &&
		!hasContentType(ct, "application/json") {
		err := fmt.Errorf("content-type must be application/json")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxBytes)
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(dst); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	return nil
}

func hasContentType(got, want string) bool {
	// match "application/json" and "application/json; charset=utf-8"
	if len(got) < len(want) {
		return false
	}
	return got[:len(want)] == want
}

// WriteText is a tiny convenience for "200 OK\n" style replies.
func WriteText(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	fmt.Fprintln(w, msg)
}
