package blocks

import "time"

func BlockTime() string {
	return time.Now().Format("Mon, 02 Jan 15:04:05 ")
}
