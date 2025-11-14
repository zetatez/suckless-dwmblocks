package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

var (
	statusbar  []string
	lastStatus string
)

var (
    SIGPLUS  syscall.Signal
    SIGMINUS syscall.Signal
)

func init() {
    initSignals()
}

func initSignals() {
    switch runtime.GOOS {
    case "openbsd":
        SIGPLUS  = syscall.Signal(int(syscall.SIGUSR1) + 1)
        SIGMINUS = syscall.Signal(int(syscall.SIGUSR1) - 1)
    case "linux":
        const defaultSIGRTMIN = 34
        SIGPLUS  = syscall.Signal(defaultSIGRTMIN)
        SIGMINUS = syscall.Signal(defaultSIGRTMIN + 1)
    default:
        SIGPLUS  = syscall.SIGUSR1
        SIGMINUS = syscall.SIGUSR2
    }
}

func main() {
	statusbar = make([]string, len(Blocks))

	sigCh := make(chan os.Signal, 10)
	SetupSignalNotifications(sigCh)

	GetCmdsAtTime(-1)
	UpdateStatus()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	var i int64 = 0
	for {
		select {
		case <-ticker.C:
			GetCmdsAtTime(i)
			UpdateStatus()
			i = (i + 1) % 86400
		case s := <-sigCh:
			switch s {
			case syscall.SIGTERM, syscall.SIGINT:
				return
			default:
				if sig, ok := s.(syscall.Signal); ok {
					off := int(sig) - int(SIGPLUS)
					if off < 0 {
						off = int(sig) - int(SIGMINUS)
					}
					if off >= 0 {
						GetSigCmds(uint(off))
						UpdateStatus()
					}
				}
			}
		}
	}
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
	if b.Func != nil {
		return b.Icon + b.Func()
	}
	if b.Command != "" {
		out, err := RunCmd(b.Command, time.Second)
		if err == nil {
			return b.Icon + out
		}
	}
	return b.Icon
}

func GetCmdsAtTime(t int64) {
    var wg sync.WaitGroup
    for i, b := range Blocks {
        if (b.Interval != 0 && t%b.Interval == 0) || t == -1 {
            wg.Add(1)
            go func(i int, b Block) {
                defer wg.Done()
                statusbar[i] = GetCmd(b)
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
	return exec.Command("xsetroot", "-name", s).Run()
}
