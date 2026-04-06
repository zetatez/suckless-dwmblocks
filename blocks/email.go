package blocks

import (
	"fmt"
	"os"
	"path"
	"strings"
)

var emailIcons = map[string]string{
	"new-email": "󱃚",
	"empty":     "󰇯",
}

func BlockEmail() string {
	cfg := getConfig()
	count, err := CountMaildirNew(cfg.EmailDir)
	if err != nil {
		return "?"
	}
	if count == 0 {
		return emailIcons["empty"]
	}
	return fmt.Sprintf("%s %d", emailIcons["new-email"], count)
}

func CountMaildirNew(maildir string) (int, error) {
	maildir = expandHome(maildir)
	newPath := path.Join(maildir, "new")
	entries, err := os.ReadDir(newPath)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		count++
	}
	return count, nil
}

func expandHome(p string) string {
	if strings.HasPrefix(p, "~/") {
		return path.Join(os.Getenv("HOME"), p[2:])
	}
	if strings.HasPrefix(p, "~") {
		return path.Join(os.Getenv("HOME"), p[1:])
	}
	if !strings.HasPrefix(p, "/") {
		return path.Join(os.Getenv("HOME"), p)
	}
	return p
}
