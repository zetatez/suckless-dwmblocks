package main

import (
	"dwmblocks/blocks"
)

var Delim = "  "

type Block struct {
	Interval int64
	Signal   uint
	Icon     string
	Func     func() string
	Command  string
}

var Blocks = []Block{
	{Interval: 3, Signal: 2, Icon: "", Func: blocks.BlockProcs},
	{Interval: 180, Signal: 3, Icon: "", Func: blocks.BlockEmail},
	{Interval: 180, Signal: 9, Icon: "", Func: blocks.BlockWeather},
	{Interval: 1, Signal: 4, Icon: "", Func: blocks.BlockVol},
	{Interval: 1, Signal: 5, Icon: "", Func: blocks.BlockMicro},
	{Interval: 1, Signal: 10, Icon: "", Func: blocks.BlockScreenLight},
	{Interval: 3, Signal: 7, Icon: "", Func: blocks.BlockNet},
	{Interval: 1800, Signal: 11, Icon: "", Func: blocks.BlockDisk},
	{Interval: 30, Signal: 6, Icon: "", Func: blocks.BlockCPUTemp},
	{Interval: 7, Signal: 11, Icon: "", Func: blocks.BlockCPU},
	{Interval: 7, Signal: 12, Icon: "", Func: blocks.BlockRAM},
	{Interval: 2, Signal: 13, Icon: "", Func: blocks.BlockBattery},
	{Interval: 1, Signal: 14, Icon: "", Func: blocks.BlockTime},
}
