package blocks

import (
	"bytes"
	"io"
	"net/http"
	"time"
)

var weatherClient = &http.Client{Timeout: 5 * time.Second}

func BlockWeather() string {
	// %c = 天气符号, %t = 温度
	resp, err := weatherClient.Get("https://wttr.in/?format=%c%t")
	if err != nil {
		return "N--"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "--"
	}

	return string(bytes.TrimSpace(body))
}
