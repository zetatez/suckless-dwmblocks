package main

import (
	"context"
	"fmt"
	"math"
	"os/exec"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"dwmblocks/api"
	"dwmblocks/svc"
)

type Block struct {
	Interval time.Duration
	Func     func() string
	nextRun  time.Time
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
	srv := svc.New(SVCAddr).Handle(api.Notify()...)
	if err := srv.Start(); err != nil {
		fmt.Printf("svc: %s\n", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	statusbar = make([]string, len(Blocks))

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
	const minWake = 100 * time.Millisecond
	now := time.Now()
	minDur := time.Duration(math.MaxInt64)
	found := false
	for i := range Blocks {
		b := &Blocks[i]
		if b.Interval <= 0 {
			continue
		}
		d := b.nextRun.Sub(now)
		if d < minDur {
			minDur = d
			found = true
		}
	}
	if !found {
		return time.Second
	}
	if minDur < minWake {
		return minWake
	}
	return minDur
}

func runBlock(b Block) (val string, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("block panic: ", r)
		}
	}()
	if b.Func != nil {
		val = b.Func()
		ok = true
	}
	return
}

func runBlockWithTimeout(block Block, timeout time.Duration) (string, bool) {
	if timeout <= 0 {
		timeout = 10 * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ch := make(chan string, 1)
	go func() {
		val, ok := runBlock(block)
		if ok {
			ch <- val
		}
	}()

	select {
	case val := <-ch:
		return val, true
	case <-ctx.Done():
		return "", false
	}
}

func RunOnceAllBlocks() {
	var wg sync.WaitGroup
	for i := range Blocks {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			block := Blocks[idx]
			timeout := block.Interval
			if timeout <= 0 {
				timeout = 10 * time.Second
			}
			if val, ok := runBlockWithTimeout(block, timeout); ok {
				statusbar[idx] = val
			}
		}(i)
	}
	wg.Wait()
}

func RunDueBlocks() {
	var wg sync.WaitGroup
	now := time.Now()
	for i := range Blocks {
		b := &Blocks[i]
		if b.Interval <= 0 || now.Before(b.nextRun) {
			continue
		}

		wg.Add(1)
		go func(idx int, block Block) {
			defer wg.Done()
			if val, ok := runBlockWithTimeout(block, block.Interval); ok {
				statusbar[idx] = val
			}
		}(i, *b)
		// Schedule next run from completion time (set after wg.Wait) below.
		b.nextRun = now.Add(b.Interval)
	}
	wg.Wait()
	// If a block took longer than its interval, push nextRun forward so we
	// don't busy-loop catching up.
	done := time.Now()
	for i := range Blocks {
		b := &Blocks[i]
		if b.Interval <= 0 {
			continue
		}
		if !done.Before(b.nextRun) {
			b.nextRun = done.Add(b.Interval)
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
