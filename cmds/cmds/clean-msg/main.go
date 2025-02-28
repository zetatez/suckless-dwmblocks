package main

import (
	"cmds/sugar"
)

const (
	MsgPath = "/tmp/.msg"
)

func main() {
	sugar.CleanMsg(MsgPath)
}
