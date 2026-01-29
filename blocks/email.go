package blocks

import (
	"fmt"
	"os"
	"path"
)

const EmailDir = ".mail/inbox"

var emailIcons = map[string]string{
	"new-email": "󱃚",
	"empty":     "󰇯",
}

func BlockEmail() string {
	count, err := CountMaildirNew(EmailDir)
	if err != nil {
		return "?"
	}
	if count == 0 {
		return emailIcons["empty"]
	}
	return fmt.Sprintf("%s %d", emailIcons["new-email"], count)
}

func CountMaildirNew(maildir string) (int, error) {
	newPath := path.Join(os.Getenv("HOME"), maildir, "new")
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
