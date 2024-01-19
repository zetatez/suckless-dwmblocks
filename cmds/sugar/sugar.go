package sugar

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

func GetWeather() (weather string, err error) {
	resp, err := http.Get("https://v2.wttr.in")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}
	slice1 := strings.Split(doc.Find("pre").Text(), "\n")
	if len(slice1) == 0 {
		return "", fmt.Errorf("no weather found")
	}
	slice2 := strings.Split(slice1[0], ",")
	if len(slice2) != 5 {
		return "", fmt.Errorf("no weather found")
	}
	temp, wind := strings.TrimSpace(slice2[1]), strings.TrimSpace(slice2[3])
	weather = fmt.Sprintf("Temp: %s, Wind: %s", temp, wind)
	return weather, nil
}

func GetClock() (clock string) {
	return time.Now().Format("Mon, Jan/02 15:04:05 ")
}

func GetCpuPercent() (percent float64, err error) {
	percents, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	percent = percents[0]
	return percent, nil
}

func GetCpuTemp() (temp float64, err error) {
	return temp, nil
}

func GetMemPercent() (percent float64, err error) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	percent = stat.UsedPercent
	return percent, nil
}

func GetDiskPercent() (percent float64, err error) {
	parts, err := disk.Partitions(false)
	if err != nil {
		return 0, err
	}
	diskInfo, err := disk.Usage(parts[0].Mountpoint)
	if err != nil {
		return 0, err
	}
	percent = diskInfo.UsedPercent
	return percent, nil
}

func GetBattery(batteryPath string) (capacity float64, status string, err error) {
	capacityByte, err := os.ReadFile(path.Join(batteryPath, "capacity"))
	if err != nil {
		return 0, "", err
	}
	capacity, err = strconv.ParseFloat(strings.TrimSpace(string(capacityByte)), 64)
	if err != nil {
		return 0, "", err
	}
	statusByte, err := os.ReadFile(path.Join(batteryPath, "status"))
	if err != nil {
		return 0, "", err
	}
	status = strings.TrimSpace(string(statusByte))
	return capacity, status, err
}

func GetNet(netPath string) (operstate string, err error) {
	operstateStr, err := os.ReadFile(
		path.Join(netPath, "operstate"),
	)
	if err != nil {
		return "", err
	}
	operstate = strings.TrimSpace(string(operstateStr))
	return operstate, err
}

type Email struct {
	From    string
	Date    string
	Subject string
}

func GetEmail(emailPath string) (emails []Email, err error) {
	inboxByte, err := os.ReadFile(path.Join(os.Getenv("HOME"), emailPath))
	if err != nil {
		return emails, err
	}
	inbox := string(inboxByte)
	r := regexp.MustCompile("From: (?P<from>.*)\nMime-Version: .*\nDate: (?P<date>.*)\nSubject: (?P<subject>.*)\n")
	xs := r.FindAllStringSubmatch(inbox, -1)
	for _, x := range xs {
		if len(x) != 4 {
			continue
		}
		emails = append(
			emails,
			Email{From: x[1], Date: x[2], Subject: x[3]},
		)
	}
	return emails, err
}

func GetMsg(msgPath string) (msg string, err error) {
	msgByte, err := os.ReadFile(msgPath)
	if err != nil {
		return "", err
	}
	msg = strings.TrimSpace(string(msgByte))
	return msg, err
}

func CleanMsg(msgPath string) {
	os.WriteFile(msgPath, []byte(""), 0o644)
	return
}

func GetProcs() (procs []*process.Process, err error) {
	return process.Processes()
}

func GetScreenLight() (percent float64, err error) {
	out, err := exec.Command("light").Output()
	if err != nil {
		return 0, err
	}
	percent, err = strconv.ParseFloat(strings.TrimSpace(string(out)), 64)
	if err != nil {
		return 0, err
	}
	return percent, nil
}

func GetVolume() (offOrOn string, percent float64, err error) {
	out, err := exec.Command("amixer", "get", "Master").Output()
	if err != nil {
		return "", 0, err
	}
	if strings.Contains(string(out), "[off]") {
		offOrOn = "off"
	} else {
		offOrOn = "on"
	}
	r := regexp.MustCompile(`\[(?P<percent>\d+)%\]`)
	xs := r.FindAllStringSubmatch(string(out), -1)
	if len(xs) != 2 {
		return offOrOn, 0, fmt.Errorf("no volume found")
	}
	if len(xs[0]) != 2 {
		return offOrOn, 0, fmt.Errorf("no volume found")
	}
	left, err := strconv.ParseFloat(xs[0][1], 64)
	if err != nil {
		return offOrOn, 0, err
	}
	right, err := strconv.ParseFloat(xs[1][1], 64)
	if err != nil {
		return offOrOn, 0, err
	}
	percent = (left + right) / 2
	return offOrOn, percent, nil
}

func GetMicro() (offOrOn string, percent float64, err error) {
	out, err := exec.Command("amixer", "get", "Capture").Output()
	if err != nil {
		return "", 0, err
	}
	if strings.Contains(string(out), "[off]") {
		offOrOn = "off"
	} else {
		offOrOn = "on"
	}
	r := regexp.MustCompile(`\[(?P<percent>\d+)%\]`)
	xs := r.FindAllStringSubmatch(string(out), -1)
	if len(xs) != 2 {
		return offOrOn, 0, fmt.Errorf("get micro failed")
	}
	if len(xs[0]) != 2 {
		return offOrOn, 0, fmt.Errorf("get micro failed")
	}
	left, err := strconv.ParseFloat(xs[0][1], 64)
	if err != nil {
		return offOrOn, 0, err
	}
	right, err := strconv.ParseFloat(xs[1][1], 64)
	if err != nil {
		return offOrOn, 0, err
	}
	percent = (left + right) / 2
	return offOrOn, percent, nil
}
