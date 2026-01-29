#!/bin/sh
set -eu

REPO="zetatez/suckless-dwmblocks"
BINARY="dwmblocks"
SERVICE_DIR="${XDG_CONFIG_HOME:-$HOME/.config}/systemd/user"

if ! command -v go >/dev/null 2>&1; then
	echo "Error: Go is not installed." >&2
	exit 1
fi

tmpdir=$(mktemp -d)
trap 'rm -rf "$tmpdir"' EXIT

echo "==> Cloning $REPO..."
git clone --depth=1 "https://github.com/$REPO.git" "$tmpdir"

echo "==> Building $BINARY..."
cd "$tmpdir"
go build -ldflags="-s -w" -o "$BINARY" .

echo "==> Installing $BINARY to /usr/local/bin..."
install -d /usr/local/bin
install -m 755 "$BINARY" /usr/local/bin/

echo "==> Setting up user service..."
install -d "$SERVICE_DIR"
install -m 644 dwmblocks.service "$SERVICE_DIR/"
systemctl --user daemon-reload
systemctl --user disable --now dwmblocks.service 2>/dev/null || true
systemctl --user enable --now dwmblocks.service

echo "==> Done: $(command -v "$BINARY")"
