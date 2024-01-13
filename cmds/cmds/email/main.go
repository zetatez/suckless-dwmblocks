package main

import (
	"fmt"
	"os"

	"cmds/sugar"
)

const (
	EmailPath = ".mail/inbox"
	MsgPath   = "/tmp/.msg"
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
		msg := ""
		for _, email := range emails {
			msg += fmt.Sprintf("Sub: %s; Fro: %s.\n", email.Subject, email.From)
		}
		os.WriteFile(MsgPath, []byte(msg), 0o644)
	default:
		str = inboxIcons["empty"]
	}
	return str
}
