package blocks

import (
	"strings"
	"sync"
	"time"
)

func BlockNotify() string {
	return GetNotify().Render()
}

const (
	NotifyScrollWidth = 52
	NotifyScrollPad   = 5
	NotifySuffix      = " | "
	NotifyDefaultTTL  = 3 * time.Second
	NotifyMinTTL      = 1 * time.Second
	NotifyMaxTTL      = 60 * time.Second
	NotifyQueueCap    = 32
)

type notifyItem struct {
	msg string
	ttl time.Duration
}

// NotifyCenter is the scrolling-message broker behind BlockNotify and the
// /notify HTTP routes. All fields are guarded by mu.
type NotifyCenter struct {
	mu       sync.Mutex
	queue    []notifyItem
	current  string
	runes    []rune    // padded buffer for scrolling
	offset   int       // scroll position in runes
	expireAt time.Time // zero ⇒ no current item
}

var (
	notifyCenterDefaultOnce sync.Once
	notifyCenterDefault     *NotifyCenter
)

// GetNotify returns the process-wide NotifyCenter singleton.
func GetNotify() *NotifyCenter {
	notifyCenterDefaultOnce.Do(func() {
		notifyCenterDefault = &NotifyCenter{}
	})
	return notifyCenterDefault
}

// Enqueue appends a message; shown after previously queued ones expire.
func (n *NotifyCenter) Enqueue(msg string, ttl time.Duration) {
	msg = strings.TrimSpace(msg)
	if msg == "" {
		return
	}
	ttl = clampNotifyTTL(ttl)

	n.mu.Lock()
	defer n.mu.Unlock()

	if n.expireAt.IsZero() {
		n.activateLocked(msg, ttl)
		return
	}
	if len(n.queue) >= NotifyQueueCap {
		n.queue = n.queue[1:]
	}
	n.queue = append(n.queue, notifyItem{msg: msg, ttl: ttl})
}

// Replace preempts the current message; the preempted one is re-queued at
// the head with its remaining TTL. The rest of the queue is preserved.
func (n *NotifyCenter) Replace(msg string, ttl time.Duration) {
	msg = strings.TrimSpace(msg)
	if msg == "" {
		return
	}
	ttl = clampNotifyTTL(ttl)

	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.expireAt.IsZero() && n.current != "" && time.Now().Before(n.expireAt) {
		remaining := time.Until(n.expireAt)
		if remaining < NotifyMinTTL {
			remaining = NotifyMinTTL
		}
		n.queue = append([]notifyItem{{msg: n.current, ttl: remaining}}, n.queue...)
		if len(n.queue) > NotifyQueueCap {
			n.queue = n.queue[:NotifyQueueCap]
		}
	}
	n.activateLocked(msg, ttl)
}

// Clear drops the current message and the entire queue.
func (n *NotifyCenter) Clear() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.queue = nil
	n.clearLocked()
}

// Render returns the next status-bar frame: scrolled body + suffix, or "".
func (n *NotifyCenter) Render() string {
	n.mu.Lock()
	defer n.mu.Unlock()

	if !n.expireAt.IsZero() && time.Now().After(n.expireAt) {
		if len(n.queue) > 0 {
			next := n.queue[0]
			n.queue = n.queue[1:]
			n.activateLocked(next.msg, next.ttl)
		} else {
			n.clearLocked()
		}
	}
	if n.current == "" {
		return ""
	}
	return n.frameLocked() + NotifySuffix
}

// frameLocked returns the next scroll frame. If the text fits the window,
// it's returned verbatim; otherwise the offset is advanced by one rune.
func (n *NotifyCenter) frameLocked() string {
	if len(n.runes) <= NotifyScrollWidth+NotifyScrollPad {
		return n.current
	}
	n.offset = (n.offset + 1) % len(n.runes)
	var sb strings.Builder
	sb.Grow(NotifyScrollWidth * 4)
	for i := 0; i < NotifyScrollWidth; i++ {
		sb.WriteRune(n.runes[(n.offset+i)%len(n.runes)])
	}
	return sb.String()
}

func (n *NotifyCenter) activateLocked(msg string, ttl time.Duration) {
	n.current = msg
	n.expireAt = time.Now().Add(ttl)
	n.offset = 0
	rs := []rune(msg)
	if len(rs) > NotifyScrollWidth {
		buf := make([]rune, 0, len(rs)+NotifyScrollPad)
		buf = append(buf, rs...)
		for i := 0; i < NotifyScrollPad; i++ {
			buf = append(buf, ' ')
		}
		n.runes = buf
	} else {
		n.runes = rs
	}
}

func (n *NotifyCenter) clearLocked() {
	n.current = ""
	n.runes = nil
	n.offset = 0
	n.expireAt = time.Time{}
}

func clampNotifyTTL(ttl time.Duration) time.Duration {
	switch {
	case ttl <= 0:
		return NotifyDefaultTTL
	case ttl < NotifyMinTTL:
		return NotifyMinTTL
	case ttl > NotifyMaxTTL:
		return NotifyMaxTTL
	default:
		return ttl
	}
}
