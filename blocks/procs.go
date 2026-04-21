package blocks

import (
	"os"
	"sort"
	"strings"
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
	procs, err := os.ReadDir("/proc")
	if err != nil {
		return "?"
	}

	running := make(map[string]struct{})
	for _, p := range procs {
		if !p.IsDir() {
			continue
		}
		name := getProcName(p.Name())
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

func getProcName(pid string) string {
	data, err := os.ReadFile("/proc/" + pid + "/comm")
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data))
}
