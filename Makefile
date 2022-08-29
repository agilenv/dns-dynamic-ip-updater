.PHONY run:
run:
	go run ./cmd/linkip/*

CMD=./cmd/linkip/*
DST=./release/linkip-$$GOOS-$$GOARCH
.PHONY build:
build:
	export GOOS=darwin; export GOARCH=amd64; go build -o $(DST) $(CMD)
	export GOOS=darwin; export GOARCH=arm64; go build -o $(DST) $(CMD)
	export GOOS=linux; export GOARCH=386; go build -o $(DST) $(CMD)
	export GOOS=linux; export GOARCH=amd64; go build -o $(DST) $(CMD)
	export GOOS=linux; export GOARCH=arm64; go build -o $(DST) $(CMD)
	export GOOS=windows; export GOARCH=386; go build -o $(DST).exe $(CMD)
	export GOOS=windows; export GOARCH=amd64; go build -o $(DST).exe $(CMD)