BINARY = dwmblocks
SRC = .
PREFIX = /usr/local

.PHONY: all build run clean install uninstall

all: build

build:
	@echo "==> Building $(BINARY)..."
	@go build -o $(BINARY) $(SRC)

run:
	@echo "==> Running $(SRC)..."
	@go run $(SRC)

clean:
	@echo "==> Cleaning..."
	@rm -f $(BINARY)

install: build
	@echo "==> Installing to $(PREFIX)/bin"
	@sudo mkdir -p $(PREFIX)/bin
	@sudo cp -f $(BINARY) $(PREFIX)/bin/

uninstall:
	@echo "==> Removing $(PREFIX)/bin/$(BINARY)"
	@sudo rm -f $(PREFIX)/bin/$(BINARY)
