package blocks

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
)

const EmailPath = ".mail/inbox"

func BlockEmail() string {
	icons := map[string]string{
		"new-email": "󱃚",
		"empty":     "󰇯",
	}
	emails, err := GetEmail(EmailPath)
	if err != nil {
		return "?"
	}
	if len(emails) == 0 {
		return icons["empty"]
	}
	var sb strings.Builder
	for _, email := range emails {
		sb.WriteString(fmt.Sprintf("Sub: %s; Fro: %s.\n", email.Subject, email.From))
	}
	if err := os.WriteFile(MsgPath, []byte(sb.String()), 0o644); err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf("%s: %d", icons["new-email"], len(emails))
}

type Email struct {
	From    string
	Date    string
	Subject string
}

func GetEmail(emailPath string) (emails []Email, err error) {
	msgByte, err := os.ReadFile(path.Join(os.Getenv("HOME"), emailPath))
	if err != nil {
		return emails, err
	}
	msg := string(msgByte)
	r := regexp.MustCompile("From: (?P<from>.*)\nMime-Version: .*\nDate: (?P<date>.*)\nSubject: (?P<subject>.*)\n")
	xs := r.FindAllStringSubmatch(msg, -1)
	for _, x := range xs {
		if len(x) != 4 {
			continue
		}
		emails = append(
			emails,
			Email{
				From:    x[1],
				Date:    x[2],
				Subject: x[3],
			},
		)
	}
	return emails, nil
}
