@echo off
set GOOS=windows
set GOARCH=amd64

go build -o bin\ripe-exporter-windows.exe .\cmd\ripe-exporter\main.go

echo Windows binary built: bin\ripe-exporter-windows.exe