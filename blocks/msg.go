package blocks

import (
	"os"
)

const MsgPath = "/tmp/.msg"

func BlockMsg() string {
	msgByte, err := os.ReadFile(MsgPath)
	if err != nil {
		if _, e := os.Stat(MsgPath); os.IsNotExist(e) {
			file, _ := os.Create(MsgPath)
			defer file.Close()
		}
		return err.Error()
	}
	return string(msgByte)
}
