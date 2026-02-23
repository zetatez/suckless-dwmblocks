package blocks

import (
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
	procs, err := process.Processes()
	if err != nil {
		return "?"
	}

	running := make(map[string]struct{})

	for _, p := range procs {
		name, err := p.Name()
		if err != nil {
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
