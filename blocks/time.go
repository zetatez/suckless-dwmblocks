package blocks

import "time"

func BlockTime() string {
	// 2006-01-02 15:04:05 Mon Jan/02 15:04:05 Jan/02 Mon 15:04:05 Jan/02 Mon 15:04:05
	return time.Now().Format("Jan/02 15:04:05 ")
}
