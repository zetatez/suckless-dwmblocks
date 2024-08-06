default: clean build

prepare:
	mkdir -p ./bins/

clean:
	rm -rf bins/*

build: prepare
	go mod tidy
	go build -o ./bins/battery         ./cmds/battery/main.go
	go build -o ./bins/clock           ./cmds/clock/main.go
	go build -o ./bins/cpu             ./cmds/cpu/main.go
	go build -o ./bins/cpu-temperature ./cmds/cpu-temperature/main.go
	go build -o ./bins/disk            ./cmds/disk/main.go
	go build -o ./bins/email           ./cmds/email/main.go
	go build -o ./bins/micro           ./cmds/micro/main.go
	go build -o ./bins/msg             ./cmds/msg/main.go
	go build -o ./bins/net             ./cmds/net/main.go
	go build -o ./bins/wifi-intensity  ./cmds/wifi-intensity/main.go
	go build -o ./bins/procs           ./cmds/procs/main.go
	go build -o ./bins/ram             ./cmds/ram/main.go
	go build -o ./bins/screen-light    ./cmds/screen-light/main.go
	go build -o ./bins/volume          ./cmds/volume/main.go
	go build -o ./bins/weather         ./cmds/weather/main.go
	go build -o ./bins/clean-msg       ./cmds/clean-msg/main.go

.PHONY: default

