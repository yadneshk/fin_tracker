build:
	go build -o $(GOPATH)/bin/fin_tracker main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o $(GOPATH)/bin/fin_tracker-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o $(GOPATH)/bin/fin_tracker-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o $(GOPATH)/bin/fin_tracker-freebsd-386 main.go