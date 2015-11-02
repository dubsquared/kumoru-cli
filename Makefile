default: clean build

linux-binary:
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o builds/linux/kumoru kumoru-cli.go

osx-binary:
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o builds/osx/kumoru kumoru-cli.go

build:
	GO15VENDOREXPERIMENT=1 go build -o kumoru kumoru-cli.go

clean:
	rm -f kumoru
	rm -f builds/osx/kumoru
	rm -f builds/linux/kumoru

restore:
	GO15VENDOREXPERIMENT=1 godep restore

release:
	clean restore build osx-binary linux-binary
