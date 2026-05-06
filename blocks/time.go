package blocks

import "time"

func BlockTime() string {
	return time.Now().Format("Mon, 1/2 15:04:05 ")
}
