.PHONY run:
run:
	go run ./cmd/cli/main.go

.PHONY build:
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o linkip ./cmd/cli/*