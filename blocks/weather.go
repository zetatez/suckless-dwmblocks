package blocks

import (
	"io"
	"net/http"
	"time"
)

func BlockWeather() string {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// %c = 天气符号, %t = 温度
	resp, err := client.Get("https://wttr.in/?format=%c%t")
	if err != nil {
		return "N--"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "--"
	}

	return string(body)
}
