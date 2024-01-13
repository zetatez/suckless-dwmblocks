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
	if len(msgByte) != 0 {
		str = fmt.Sprintf("[%s]", strings.TrimSpace(string(msgByte)))
	}
	return str
}
