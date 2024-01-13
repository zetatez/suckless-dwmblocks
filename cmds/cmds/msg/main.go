package main

import (
	"fmt"
	"strings"

	"cmds/sugar"
)

const (
	MsgPath = "/tmp/.msg"
)

func main() {
	fmt.Println(FormatMsg())
}

func FormatMsg() (str string) {
	msgByte, err := sugar.GetMsg(MsgPath)
	if err != nil {
		return ""
	}
	msg := strings.TrimSpace(string(msgByte))
	return msg
}
