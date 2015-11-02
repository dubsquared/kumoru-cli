default: clean build

linux-binary:
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o kumoru-lnx-cli kumoru-cli.go

build: 
	GO15VENDOREXPERIMENT=1 go build -o kumoru-cli kumoru-cli.go

clean: 
	rm -f kumoru-cli
