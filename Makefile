build:
	CGO_ENABLED=0 GOARCH=amd64 go build -o bin/go-server -ldflags "-w"
	chmod u+x go-server

image:
	docker build . -t devopsws/go-server
