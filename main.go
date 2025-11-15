package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"
)

type Block struct {
	Signal   uint
	Interval time.Duration
	Func     func() string
	Command  string

	nextRun time.Time
}

var (
	statusbar  []string
	lastStatus string
)

var (
	SIGPLUS  syscall.Signal
	SIGMINUS syscall.Signal
)

func init() {
	initBlocks()
	initSignals()
}

func initBlocks() {
	now := time.Now()
	for i := range Blocks {
		if Blocks[i].Interval > 0 {
			Blocks[i].nextRun = now
		}
	}
}

func initSignals() {
	switch runtime.GOOS {
	case "openbsd":
		SIGPLUS = syscall.Signal(int(syscall.SIGUSR1) + 1)
		SIGMINUS = syscall.Signal(int(syscall.SIGUSR1) - 1)
	case "linux":
		const defaultSIGRTMIN = 34
		SIGPLUS = syscall.Signal(defaultSIGRTMIN)
		SIGMINUS = syscall.Signal(defaultSIGRTMIN + 1)
	default:
		SIGPLUS = syscall.SIGUSR1
		SIGMINUS = syscall.SIGUSR2
	}
}

func main() {
	statusbar = make([]string, len(Blocks))

	sigCh := make(chan os.Signal, 10)
	SetupSignalNotifications(sigCh)

	RunOnceAllBlocks()
	UpdateStatus()

	for {
		next := NextWakeUpDuration()
		select {
		case <-time.After(next):
			RunDueBlocks()
			UpdateStatus()
		case s := <-sigCh:
			switch s {
			case syscall.SIGTERM, syscall.SIGINT:
				return
			default:
				if sig, ok := s.(syscall.Signal); ok {
					idx, ok := BlockIdxFromSignal(sig)
					if ok {
						GetSigCmds(uint(idx))
						UpdateStatus()
					}
				}
			}
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

func BlockIdxFromSignal(sig syscall.Signal) (int, bool) {
	if sig >= SIGPLUS {
		return int(sig - SIGPLUS), true
	}
	if sig >= SIGMINUS {
		return int(sig - SIGMINUS), true
	}
	return 0, false
}

func SetupSignalNotifications(sigc chan os.Signal) {
	signal.Notify(sigc, syscall.SIGTERM, syscall.SIGINT)
	for _, b := range Blocks {
		if b.Signal > 0 {
			signal.Notify(sigc, syscall.Signal(int(SIGMINUS)+int(b.Signal)))
		}
	}
}

func RunCmd(cmd string, timeout time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	c := exec.CommandContext(ctx, "sh", "-c", cmd)
	var out bytes.Buffer
	c.Stdout = &out
	c.Stderr = &out
	err := c.Run()
	return strings.TrimSpace(out.String()), err
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
	if b.Command != "" {
		out, err := RunCmd(b.Command, time.Second)
		if err == nil {
			return out
		}
	}
	return ""
}

func RunOnceAllBlocks() {
	var wg sync.WaitGroup
	for i := range Blocks {
		b := &Blocks[i]
		wg.Add(1)
		go func(i int, b *Block) {
			defer wg.Done()
			statusbar[i] = GetCmd(*b)
		}(i, b)
	}
	wg.Wait()
}

func RunDueBlocks() {
	now := time.Now()
	var wg sync.WaitGroup
	for i := range Blocks {
		b := &Blocks[i]
		if b.Interval > 0 && !now.Before(b.nextRun) {
			wg.Add(1)
			go func(i int, b *Block) {
				defer wg.Done()
				statusbar[i] = GetCmd(*b)
				b.nextRun = now.Add(b.Interval)
			}(i, b)
		}
	}
	wg.Wait()
}

func GetSigCmds(sig uint) {
	for i, b := range Blocks {
		if b.Signal == sig {
			statusbar[i] = GetCmd(b)
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

func write(s string) error {
	_, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	return exec.Command("xsetroot", "-name", s).Run()
}
