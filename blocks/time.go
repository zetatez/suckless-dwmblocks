package blocks

import "time"

func BlockTime() string {
	return time.Now().Format("Mon, Jan/02 15:04:05 ")
}
