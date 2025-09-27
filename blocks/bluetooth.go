package blocks

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"strings"
)

func BlockBluetoothConnectedDevices() string {
	connectedDevices, err := GetBlueToothConnectedDevices()
	if err != nil {
		return " ?"
	}
	if len(connectedDevices) == 0 {
		return ""
	}
	str := " "
	delimiter := "|"
	for _, device := range connectedDevices {
		str += fmt.Sprintf("%s%s", device.Name, delimiter)
	}
	return strings.TrimSuffix(str, delimiter)
}

type Device struct {
	Name string
	Addr string
}

func GetBlueToothConnectedDevices() ([]Device, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %v", err)
	}
	// get all objects managed by BlueZ
	manager := conn.Object("org.bluez", "/")
	var objs map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	err = manager.Call("org.freedesktop.DBus.ObjectManager.GetManagedObjects", 0).Store(&objs)
	if err != nil {
		return nil, fmt.Errorf("failed to get managed objects: %v", err)
	}
	connectedDevices := make([]Device, 0)
	for _, ifaces := range objs {
		if dev, ok := ifaces["org.bluez.Device1"]; ok {
			connected := dev["Connected"].Value().(bool)
			if connected {
				name := dev["Name"].Value().(string)
				addr := dev["Address"].Value().(string)
				connectedDevices = append(connectedDevices, Device{Name: name, Addr: addr})
			}
		}
	}
	return connectedDevices, nil
}
