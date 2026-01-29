// Package api wires domain features to HTTP routes for the svc server.
package api

import (
	"net/http"
	"strings"
	"time"

	"dwmblocks/blocks"
	"dwmblocks/svc"
)

// Notify returns the svc.Route values implementing the notify protocol.
//
//	POST   /notify   enqueue
//	PUT    /notify   replace (preempted item re-queued at head)
//	DELETE /notify   clear current + queue
//
// Body for POST/PUT must be application/json:
//
//	{"msg":"...", "ttl_seconds":30}
func Notify() []svc.Route {
	n := blocks.GetNotify()
	return []svc.Route{
		{Method: http.MethodPost, Path: "/notify", Handler: notifyWrite(n.Enqueue)},
		{Method: http.MethodPut, Path: "/notify", Handler: notifyWrite(n.Replace)},
		{Method: http.MethodDelete, Path: "/notify", Handler: notifyClear(n)},
	}
}

func notifyClear(n *blocks.NotifyCenter) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		n.Clear()
		w.WriteHeader(http.StatusNoContent)
	}
}

func notifyWrite(action func(string, time.Duration)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body struct {
			Msg        string `json:"msg"`
			TTLSeconds int    `json:"ttl_seconds"`
		}
		if err := svc.DecodeJSON(w, r, &body, 0); err != nil {
			return
		}
		msg := strings.TrimSpace(body.Msg)
		if msg == "" {
			http.Error(w, "empty msg", http.StatusBadRequest)
			return
		}
		ttl := time.Duration(body.TTLSeconds) * time.Second
		action(msg, ttl)
		svc.WriteText(w, http.StatusOK, "ok")
	}
}
