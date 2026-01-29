package blocks

type Config struct {
	BatteryPath  string
	EmailDir     string
	NetInterface string
}

var blockConfig = Config{
	BatteryPath:  "/sys/class/power_supply/BAT0",
	EmailDir:     "~/.mail/inbox",
	NetInterface: "wlan0",
}

func SetConfig(cfg Config) {
	blockConfig = cfg
}

func getConfig() Config {
	return blockConfig
}
