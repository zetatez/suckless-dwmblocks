package main

import (
	"dwmblocks/blocks"
	"time"
)

var Delim = "  "

var Blocks = []Block{
	{Signal: 7, Interval: 5 * time.Second, Func: blocks.BlockProcs},
	{Signal: 7, Interval: 15 * time.Minute, Func: blocks.BlockWeather},
	{Signal: 7, Interval: 15 * time.Minute, Func: blocks.BlockEmail},
	{Signal: 7, Interval: 2 * time.Second, Func: blocks.BlockBluetoothConnectedDevices},
	{Signal: 7, Interval: 1 * time.Second, Func: blocks.BlockVol},
	{Signal: 7, Interval: 1 * time.Second, Func: blocks.BlockMicro},
	{Signal: 7, Interval: 1 * time.Second, Func: blocks.BlockScreenLight},
	// {Signal: 7, Interval: 30 * time.Minute, Func: blocks.BlockDisk},
	{Signal: 7, Interval: 1 * time.Minute, Func: blocks.BlockCPUTemp},
	{Signal: 7, Interval: 7 * time.Second, Func: blocks.BlockCPU},
	{Signal: 7, Interval: 7 * time.Second, Func: blocks.BlockRAM},
	{Signal: 7, Interval: 1 * time.Second, Func: blocks.BlockInputMethod},
	{Signal: 7, Interval: 3 * time.Second, Func: blocks.BlockBattery},
	{Signal: 7, Interval: 3 * time.Second, Func: blocks.BlockNet},
	{Signal: 7, Interval: 1 * time.Second, Func: blocks.BlockTime},
}
