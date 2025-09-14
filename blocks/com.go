package blocks

func GetIconByPct(pct float64) (icon string) {
	icons := map[string]string{
		"5":   "",
		"10":  "",
		"25":  "󰖃",
		"50":  "󰜎",
		"75":  "󰑮",
		"100": "󱄟",
	}
	switch {
	case pct < 10:
		icon = icons["10"]
	case pct < 25:
		icon = icons["25"]
	case pct < 50:
		icon = icons["50"]
	case pct < 75:
		icon = icons["75"]
	default:
		icon = icons["100"]
	}
	return icon
}
