package blocks

import (
	"fmt"
	"sync"

	"github.com/godbus/dbus/v5"
)

var (
	fcitxOnce sync.Once
	fcitxConn *dbus.Conn
	fcitxObj  dbus.BusObject
	fcitxErr  error
)

func BlockInputMethod() string {
	inputMethod, err := GetInputMethod()
	if err != nil {
		return " ?"
	}
	return fmt.Sprintf("  %s", inputMethod)
}

func GetInputMethod() (string, error) {
	fcitxOnce.Do(func() {
		fcitxConn, fcitxErr = dbus.ConnectSessionBus()
		if fcitxErr == nil {
			fcitxObj = fcitxConn.Object("org.fcitx.Fcitx5", "/controller")
		}
	})
	if fcitxErr != nil {
		return "", fcitxErr
	}
	var inputMethod string
	if err := fcitxObj.Call("org.fcitx.Fcitx.Controller1.CurrentInputMethod", 0).Store(&inputMethod); err != nil {
		return "", err
	}
	switch inputMethod {
	case "pinyin":
		return "中文", nil
	case "keyboard-us":
		return "英文", nil
	default:
		return inputMethod, nil
	}
}
