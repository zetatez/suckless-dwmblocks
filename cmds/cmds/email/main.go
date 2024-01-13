package main

import (
	"fmt"

	"cmds/sugar"
)

const (
	EmailPath = ".mail/inbox"
)

func main() {
	fmt.Println(FormatEmail())
}

func FormatEmail() (str string) {
	emails, err := sugar.GetEmail(EmailPath)
	if err != nil {
		return ""
	}
	inboxIcons := map[string]string{
		"new-email": "📩",
		"empty":     "📨",
	}
	switch {
	case len(emails) > 0:
		str = fmt.Sprintf("%s: %d", inboxIcons["new-email"], len(emails))
	default:
		str = inboxIcons["empty"]
	}
	return str
}
