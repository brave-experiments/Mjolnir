default: test

generate:
	go run builder/main.go

test: clean generate
	go test -cover -covermode=count -coverprofile=coverage.out ./...

test-silent: clean-build generate
	go test -cover -covermode=count -coverprofile=coverage.out ./... > dist/${CLI_VERSION}/unit.log

clean-build: clean
	rm -rf dist/${CLI_VERSION}/unix
	mkdir -p dist/${CLI_VERSION}/unix

clean-build-mac: clean
	rm -rf dist/${CLI_VERSION}/osx
	mkdir -p dist/${CLI_VERSION}/osx

clean:
	rm -rf coverage.out

build: build-unix build-mac
	ls -la dist/${CLI_VERSION}/

build-unix: generate clean-build
	CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/unix/apollo

build-mac: generate clean-build-mac
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/osx/apollo
	ls -la dist/${CLI_VERSION}/osx/apollo

test-and-build: clean clean-build generate
	go test -cover -covermode=count -coverprofile=coverage.out ./...
	GOPROXY=https://proxy.golang.org CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/unix/apollo
