package blocks

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/shirou/gopsutil/process"
)

var concernedProcsIcon = map[string]string{
	"flameshot":           "",
	"vim":                 "",
	"nvim":                "",
	"subl":                "",
	"ffmpeg":              "",
	"inkscape":            "",
	"krita":               "",
	"netease-cloud-music": "󰝚",
	"obsidian":            "󱓩",
	"zoom":                "󱐒",
	"xournalpp":           "󰽉",
	"zathura":             "",
	"dockerd":             "",
	"chrome":              "󰊯",
	"qutebrowser":         "",
	"clash":               "🌐",
}

func BlockProcs() string {
	pids, err := process.Pids()
	if err != nil {
		return "?"
	}

	running := make(map[string]struct{})
	for _, pid := range pids {
		name := getProcName(pid)
		if name == "" {
			continue
		}
		if _, ok := concernedProcsIcon[name]; ok {
			running[name] = struct{}{}
			if len(running) == len(concernedProcsIcon) {
				break
			}
		}
	}

	icons := make([]string, 0, len(running))
	for proc := range running {
		icons = append(icons, concernedProcsIcon[proc])
	}
	sort.Strings(icons)

	return "< " + strings.Join(icons, " ") + " >"
}

func getProcName(pid int32) string {
	data, err := os.ReadFile(fmt.Sprintf("/proc/%d/comm", pid))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}
