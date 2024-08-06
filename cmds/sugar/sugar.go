package sugar

import (
	"fmt"
	"net"
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
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

func Notify(msg ...interface{}) {
	NewExecService().RunScriptShell(fmt.Sprintf("notify-send '%v'", msg))
}

func GetWeather() (temp, wind string, err error) {
	resp, err := http.Get("https://v2.wttr.in")
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", "", err
	}
	slice1 := strings.Split(doc.Find("pre").Text(), "\n")
	if len(slice1) == 0 {
		return "", "", fmt.Errorf("no weather found")
	}
	slice2 := strings.Split(slice1[0], ",")
	if len(slice2) != 5 {
		return "", "", fmt.Errorf("no weather found")
	}
	temp, wind = strings.TrimSpace(slice2[1]), strings.TrimSpace(slice2[3])
	return temp, wind, nil
}

func GetClock() (clock string) {
	// return time.Now().Format("Mon, Jan/02 15:04:05 ")
	return time.Now().Format("Jan/02 Mon 15:04:05 ")
}

func GetCpuTemperature() (avgTemerature float64, err error) {
	sensors, err := host.SensorsTemperatures()
	if err != nil {
		return 0, err
	}
	ct, sum := 0.0, 0.0
	for _, sensor := range sensors {
		if strings.HasPrefix(sensor.SensorKey, "coretemp_core") && strings.HasSuffix(sensor.SensorKey, "_input") {
			ct++
			sum += sensor.Temperature
		}
	}
	return sum / ct, nil
}

func GetCpuPercent() (percent float64, err error) {
	percents, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	percent = percents[0]
	return percent, nil
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
	if len(xs) != 1 {
		return offOrOn, 0, fmt.Errorf("no volume found")
	}
	percent, err = strconv.ParseFloat(xs[0][1], 64)
	if err != nil {
		return offOrOn, 0, err
	}
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

func GetLocalIpv4ByInterfaceName(interfaceName string) (addr string, err error) {
	i, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", err
	}
	addrs, err := i.Addrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok || ipNet.IP.To4() == nil {
			continue
		}
		return ipNet.IP.To4().String(), nil
	}
	return "", fmt.Errorf("interface %s don't have an ipv4 addr", interfaceName)
}

func GetActiveWifi() (ssid string, signal int) {
	stdout, _, err := NewExecService().RunScriptShell("nmcli -t -f ACTIVE,SSID,SIGNAL device wifi")
	if err != nil {
		return "", 0
	}
	lines := strings.Split(string(stdout), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) == 3 && fields[0] == "yes" {
			ssid = fields[1]
			signal, _ = strconv.Atoi(fields[2])
			return ssid, signal
		}
	}
	return "", 0
}
