BUILD_FLAGS :=  "-X main.BuildStamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.GitVersion=`git rev-parse HEAD`"

default: clean build

linux-binary:
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o builds/linux/kumoru kumoru-cli.go

osx-binary:
	GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -ldflags $(BUILD_FLAGS) -o builds/osx/kumoru kumoru-cli.go

build:
	GO15VENDOREXPERIMENT=1 go build -a -ldflags $(BUILD_FLAGS) -o kumoru kumoru-cli.go

install: build mv-bin

mv-bin:
	cp kumoru ${GOPATH}/bin/

clean:
	rm -f kumoru
	rm -f builds/osx/kumoru
	rm -f builds/linux/kumoru

restore:
	GO15VENDOREXPERIMENT=1 godep restore

depsave:
	rm -f Godeps/Godeps.json
	GO15VENDOREXPERIMENT=1 godep save

test:
	GO15VENDOREXPERIMENT=1 go test -cover ./...

release: clean restore test osx-binary linux-binary
