package main

import (
	"fmt"
	"sort"
	"strings"

	"cmds/sugar"
)

func main() {
	fmt.Println(FormatProcs())
}

func FormatProcs() (str string) {
	concernedProcsIcon := map[string]string{
		"chrome":              "󰊭",
		"lazygit":             "󰊢",
		"clash":               "🌏",
		"ffmpeg":              "",
		"flameshot":           "",
		"inkscape":            "",
		"yazi":                "",
		"krita":               "",
		"ncmpcpp":             "󰝚",
		"netease-cloud-music": "󰝚",
		"obsidian":            "󱓩",
		"screenkey":           "",
		"subl":                "",
		"vim":                 "",
		"wechat-uos":          "󰘑",
		"wemeet":              "󱋒",
		"xournalpp":           "󰽉",
		"zathura":             "",
		"zoom":                "󱐒",
		"dockerd":             "",
		// "python":              "",
		// "julia":               "",
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
	procIconList := []string{}
	for proc := range runningConcernedProcs {
		procIconList = append(procIconList, concernedProcsIcon[proc])
	}
	sort.Strings(procIconList)
	str = "< " + strings.Join(procIconList, " ") + " >"
	return str
}
