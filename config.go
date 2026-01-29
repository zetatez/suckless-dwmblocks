package main

import (
	"time"

	"dwmblocks/blocks"
)

var Delim = " "

// SVCAddr is the HTTP control endpoint listen address.
// Default listens on all interfaces; firewall externally if not desired.
// Example: curl -X POST -d 'hello' http://127.0.0.1:8765/notify
var SVCAddr = "0.0.0.0:8765"

var Blocks = []Block{
	{Interval: 200 * time.Millisecond, Func: blocks.BlockNotify},
	{Interval: 3 * time.Second, Func: blocks.BlockBluetoothConnectedDevices},
	{Interval: 1 * time.Second, Func: blocks.BlockVol},
	{Interval: 1 * time.Second, Func: blocks.BlockMicro},
	{Interval: 1 * time.Second, Func: blocks.BlockScreenLight},
	{Interval: 3 * time.Second, Func: blocks.BlockCPU},
	{Interval: 3 * time.Second, Func: blocks.BlockMem},
	{Interval: 3 * time.Second, Func: blocks.BlockBattery},
	{Interval: 3 * time.Second, Func: blocks.BlockNet},
	{Interval: 1 * time.Second, Func: blocks.BlockInputMethod},
	{Interval: 1 * time.Hour, Func: blocks.BlockWeather},
	{Interval: 1 * time.Second, Func: blocks.BlockTime},
}
