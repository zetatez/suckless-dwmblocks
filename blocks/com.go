package blocks

func GetIconByPct(pct float64) (icon string) {
	switch {
	case pct < 10:
		icon = ""
	case pct < 25:
		icon = "󰖃"
	case pct < 50:
		icon = "󰜎"
	case pct < 75:
		icon = "󰑮"
	default:
		icon = "󱄟"
	}
	return icon
}
