package main

import (
	"context"
	"fmt"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Block struct {
	Interval time.Duration
	Func     func() string

	nextRun time.Time
}

var (
	statusbar  []string
	lastStatus string
)

func init() {
	initBlocks()
}

func initBlocks() {
	now := time.Now()
	for i := range Blocks {
		if Blocks[i].Interval > 0 {
			Blocks[i].nextRun = now
		}
	}
}

func main() {
	statusbar = make([]string, len(Blocks))
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	RunOnceAllBlocks()
	UpdateStatus()

	timer := time.NewTimer(0)
	defer timer.Stop()

	for {
		next := NextWakeUpDuration()
		if !timer.Stop() {
			select {
			case <-timer.C:
			default:
			}
		}
		timer.Reset(next)

		select {
		case <-timer.C:
			RunDueBlocks()
			UpdateStatus()
		case <-ctx.Done():
			return
		}
	}
}

func NextWakeUpDuration() time.Duration {
	now := time.Now()
	next := time.Duration(1 << 62)
	for _, b := range Blocks {
		if b.Interval <= 0 {
			continue
		}
		if b.nextRun.After(now) {
			d := b.nextRun.Sub(now)
			if d < next {
				next = d
			}
		} else {
			return 0
		}
	}
	if next == time.Duration(1<<62) {
		return time.Second
	}
	return next
}

func GetCmd(b Block) string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("block panic: ", r)
		}
	}()
	if b.Func != nil {
		return b.Func()
	}
	return ""
}

func RunOnceAllBlocks() {
	for i := range Blocks {
		statusbar[i] = GetCmd(Blocks[i])
	}
}

func RunDueBlocks() {
	now := time.Now()
	for i := range Blocks {
		b := &Blocks[i]
		if b.Interval > 0 && !now.Before(b.nextRun) {
			statusbar[i] = GetCmd(*b)
			b.nextRun = now.Add(b.Interval)
		}
	}
}

func BuildStatus() string {
	var sb strings.Builder
	for _, s := range statusbar {
		if s != "" {
			if sb.Len() > 0 {
				sb.WriteString(Delim)
			}
			sb.WriteString(s)
		}
	}
	return sb.String()
}

func UpdateStatus() {
	status := BuildStatus()
	if status != lastStatus {
		if err := write(status); err != nil {
			fmt.Printf("Error writing status: %s", err)
		}
		lastStatus = status
	}
}

func Notify(msg ...any) {
	msgStr := strings.TrimSpace(fmt.Sprint(msg...))
	if msgStr == "" {
		return
	}
	exec.Command("notify-send", msgStr).Run()
}
