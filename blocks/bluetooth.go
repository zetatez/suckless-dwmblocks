package blocks

import (
	"fmt"
	"github.com/godbus/dbus/v5"
	"strings"
	"sync"
)

var (
	bluezOnce sync.Once
	bluezConn *dbus.Conn
	bluezErr  error
)

func BlockBluetoothConnectedDevices() string {
	connectedDevices, err := GetBlueToothConnectedDevices()
	if err != nil {
		return " ?"
	}
	if len(connectedDevices) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(" ")
	for i, device := range connectedDevices {
		if i > 0 {
			sb.WriteByte('|')
		}
		sb.WriteString(device.Name)
	}
	return sb.String()
}

type Device struct {
	Name string
	Addr string
}

func GetBlueToothConnectedDevices() ([]Device, error) {
	bluezOnce.Do(func() {
		bluezConn, bluezErr = dbus.SystemBus()
	})
	if bluezErr != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %v", bluezErr)
	}
	// get all objects managed by BlueZ
	manager := bluezConn.Object("org.bluez", "/")
	var objs map[dbus.ObjectPath]map[string]map[string]dbus.Variant
	if err := manager.Call("org.freedesktop.DBus.ObjectManager.GetManagedObjects", 0).Store(&objs); err != nil {
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
