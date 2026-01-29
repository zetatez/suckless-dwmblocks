BINARY = dwmblocks
PREFIX = /usr/local/bin

.PHONY: all build run clean install uninstall

all: build

build:
	go build -ldflags="-s -w" -o $(BINARY) .

run:
	go run .

clean:
	rm -f $(BINARY)

SERVICE_DIR = $(or $(XDG_CONFIG_HOME),$(HOME)/.config)/systemd/user

install: build
	sudo install -d $(PREFIX)
	sudo install -m 755 $(BINARY) $(PREFIX)/
	install -d $(SERVICE_DIR)
	install -m 644 dwmblocks.service $(SERVICE_DIR)/
	systemctl --user daemon-reload
	systemctl --user disable --now dwmblocks.service 2>/dev/null || true
	systemctl --user enable --now dwmblocks.service

uninstall:
	sudo rm -f $(PREFIX)/$(BINARY)
	rm -f $(SERVICE_DIR)/dwmblocks.service
	systemctl --user daemon-reload
