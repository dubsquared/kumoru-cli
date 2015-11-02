default: clean build

linux-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kumoru-lnx-cli kumoru-cli.go

build: 
	go build -o kumoru-cli kumoru-cli.go

clean: 
	rm -f kumoru-cli
