package main

import (
	"time"

	"dwmblocks/blocks"
)

func init() {
	blocks.SetConfig(
		blocks.Config{
			BatteryPath:  "/sys/class/power_supply/BAT0",
			EmailDir:     "~/.mail/inbox",
			NetInterface: "wlan0",
		},
	)
}

var Delim = "  "

var Blocks = []Block{
	// {Interval: 400 * time.Millisecond, Func: blocks.BlockNews},
	{Interval: 15 * time.Second, Func: blocks.BlockProcs},
	// {Interval: 15 * time.Minute, Func: blocks.BlockWeather},
	// {Interval: 15 * time.Minute, Func: blocks.BlockEmail},
	// {Interval: 3 * time.Second, Func: blocks.BlockBluetoothConnectedDevices},
	{Interval: 1 * time.Second, Func: blocks.BlockVol},
	{Interval: 1 * time.Second, Func: blocks.BlockMicro},
	{Interval: 1 * time.Second, Func: blocks.BlockScreenLight},
	// {Interval: 30 * time.Minute, Func: blocks.BlockDisk},
	// {Interval: 1 * time.Minute, Func: blocks.BlockCPUTemp},
	{Interval: 3 * time.Second, Func: blocks.BlockCPU},
	{Interval: 3 * time.Second, Func: blocks.BlockMem},
	{Interval: 1 * time.Second, Func: blocks.BlockInputMethod},
	{Interval: 3 * time.Second, Func: blocks.BlockBattery},
	{Interval: 3 * time.Second, Func: blocks.BlockNet},
	{Interval: 1 * time.Second, Func: blocks.BlockTime},
}
