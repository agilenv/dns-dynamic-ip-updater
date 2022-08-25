.PHONY run:
run:
	go run ./cmd/ddip_updater/main.go

.PHONY build:
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ddip_updater ./cmd/ddip_updater/main.go