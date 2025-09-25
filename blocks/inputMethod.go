package blocks

import (
	"fmt"
	"github.com/godbus/dbus/v5"
)

func BlockInputMethod() string {
	inputMethod, err := GetInputMethod()
	if err != nil {
		return "󰌌 ?"
	}
	return fmt.Sprintf("󰌌 %s", inputMethod)
}

func GetInputMethod() (string, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return "", err
	}
	defer conn.Close()
	obj := conn.Object("org.fcitx.Fcitx5", "/controller")
	var inputMethod string
	if err = obj.Call("org.fcitx.Fcitx.Controller1.CurrentInputMethod", 0).Store(&inputMethod); err != nil {
		return "", err
	}
	switch inputMethod {
	case "pinyin":
		return "CH", nil
	case "keyboard-us":
		return "EN", nil
	default:
		return inputMethod, nil
	}
}
