package blocks

import (
	"sort"
	"strings"

	"github.com/shirou/gopsutil/process"
)

var concernedProcsIcon = map[string]string{
	"flameshot":           "",
	"vim":                 "",
	"subl":                "",
	"ffmpeg":              "",
	"inkscape":            "",
	"krita":               "",
	"ncmpcpp":             "󰝚",
	"netease-cloud-music": "󰝚",
	"obsidian":            "󱓩",
	"wechat-uos":          "󰘑",
	"wemeet":              "󱋒",
	"zoom":                "󱐒",
	"xournalpp":           "󰽉",
	"zathura":             "",
	"dockerd":             "",
	"chrome":              "󰊭",
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
