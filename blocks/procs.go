package blocks

import (
	"sort"
	"strings"

	"github.com/shirou/gopsutil/process"
)

func BlockProcs() string {
	concernedProcsIcon := map[string]string{
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
		cmdline, err := p.Cmdline()
		if err != nil {
			cmdline = ""
		}

		for proc := range concernedProcsIcon {
			if name == proc || strings.Contains(cmdline, proc) {
				running[proc] = struct{}{}
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
