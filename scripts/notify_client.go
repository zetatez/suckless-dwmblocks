// notify_client.go — exercise all three /notify endpoints via a Client service.
//
// Run:    go run scripts/notify_client.go
// Build:  go build -o /tmp/notify_client scripts/notify_client.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const notifyURL = "http://127.0.0.1:8765/notify"

type Client struct {
	HTTP *http.Client
}

func NewClient() *Client {
	return &Client{HTTP: &http.Client{Timeout: 3 * time.Second}}
}

type notifyBody struct {
	Msg        string `json:"msg"`
	TTLSeconds int    `json:"ttl_seconds"`
}

// Post enqueues a message (shown after previous messages expire).
func (c *Client) Post(msg string, ttl time.Duration) error {
	buf, err := json.Marshal(notifyBody{Msg: msg, TTLSeconds: int(ttl.Seconds())})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, notifyURL, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(req)
}

// Put preempts the current message; the preempted one is re-queued at the head.
func (c *Client) Put(msg string, ttl time.Duration) error {
	buf, err := json.Marshal(notifyBody{Msg: msg, TTLSeconds: int(ttl.Seconds())})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPut, notifyURL, bytes.NewReader(buf))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return c.do(req)
}

// Delete clears the current message and the entire queue.
func (c *Client) Delete() error {
	req, err := http.NewRequest(http.MethodDelete, notifyURL, nil)
	if err != nil {
		return err
	}
	return c.do(req)
}

func (c *Client) do(req *http.Request) error {
	resp, err := c.HTTP.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, _ := io.ReadAll(resp.Body)
	fmt.Printf("%s %s → %s %s\n", req.Method, req.URL.Path, resp.Status, bytes.TrimSpace(out))
	if resp.StatusCode >= 300 {
		return fmt.Errorf("status %d", resp.StatusCode)
	}
	return nil
}

func main() {
	c := NewClient()

	must(c.Post("queued #1", 5*time.Second))
	must(c.Post("queued #2", 5*time.Second))
	must(c.Put("URGENT preempt", 4*time.Second))
	time.Sleep(1 * time.Second)
	must(c.Delete())
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
