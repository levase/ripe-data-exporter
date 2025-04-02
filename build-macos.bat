@echo off
set GOOS=darwin
set GOARCH=amd64

go build -o bin/ripe-exporter-macos ./cmd/ripe-exporter/main.go

echo "macOS binary built: bin/ripe-exporter-macos"