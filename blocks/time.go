package blocks

import "time"

func BlockTime() string {
	// 2006-01-02 15:04:05 Mon Jan/02
	return time.Now().Format("Mon. Jan/02 15:04:05 ")
}
