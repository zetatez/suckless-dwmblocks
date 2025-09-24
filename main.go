package main

import (
	"bytes"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
	"fmt"
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
	if runtime.GOOS == "openbsd" {
		SIGPLUS = syscall.Signal(int(syscall.SIGUSR1) + 1)
		SIGMINUS = syscall.Signal(int(syscall.SIGUSR1) - 1)
	} else {
		const linuxSIGRTMIN = 34
		SIGPLUS = syscall.Signal(linuxSIGRTMIN)
		SIGMINUS = syscall.Signal(linuxSIGRTMIN)
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

func RunCmd(cmd string) (string, error) {
	c := exec.Command("sh", "-c", cmd)
	var out bytes.Buffer
	c.Stdout = &out
	return strings.TrimSpace(out.String()), c.Run()
}

func GetCmd(b Block) string {
	if b.Func != nil {
		return b.Icon + b.Func()
	}
	if b.Command != "" {
		out, err := RunCmd(b.Command)
		if err == nil {
			return b.Icon + out
		}
	}
	return b.Icon
}

func GetCmdsAtTime(t int64) {
	for i, b := range Blocks {
		if (b.Interval != 0 && t%b.Interval == 0) || t == -1 {
			statusbar[i] = GetCmd(b)
		}
	}
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
