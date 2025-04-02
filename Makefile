.SILENT:

# Сборка под Windows
build-win:
	@set GOOS=windows
	@set GOARCH=amd64
	@go build -o bin\ripe-exporter-windows.exe ./cmd/ripe-exporter/main.go

# Сборка под macOS (Intel x86_64)
build-mac:
	@set GOOS=darwin
	@set GOARCH=amd64
	@go build -o bin/ripe-exporter-mac-amd64 ./cmd/ripe-exporter/main.go

# Сборка под macOS (Apple arm64)
build-mac-arm:
	@set GOOS=darwin
	@set GOARCH=arm64
	@go build -o bin/ripe-exporter-mac-arm64 ./cmd/ripe-exporter/main.go