package main

import (
	"fmt"
	"strings"

	"cmds/sugar"
)

const (
	BatteryPath = "/sys/class/power_supply/BAT0"
	NetPath     = "/sys/class/net/wlan0"
	EmailPath   = ".mail/inbox"
	MsgPath     = ".msg"
)

func main() {
	fmt.Println(FormatProcs())
}

func FormatProcs() (str string) {
	concernedProcsIcon := map[string]string{
		"chrome":              "",
		"wechat-uos":          "󰘑",
		"msedge":              "󰌀",
		"flameshot":           "",
		"clash":               "󰟾",
		"ffmpeg":              "󰻃",
		"ncmpcpp":             "󰝚",
		"netease-cloud-music": "󰝚",
		"vim":                 "",
		"zathura":             "",
	}
	procs, err := sugar.GetProcs()
	if err != nil {
		return ""
	}
	runningConcernedProcs := make(map[string]bool)
	for _, p := range procs {
		for concernedProc := range concernedProcsIcon {
			name, err := p.Name()
			if err != nil {
				continue
			}
			if name == concernedProc {
				runningConcernedProcs[concernedProc] = true
			} else {
				cmdline, err := p.Cmdline()
				if err != nil {
					continue
				}
				if strings.Contains(cmdline, concernedProc) {
					runningConcernedProcs[concernedProc] = true
				}
			}
		}
	}
	for proc := range runningConcernedProcs {
		str = fmt.Sprintf("%s %s", str, concernedProcsIcon[proc])
	}
	str += "|"
	return str
}
