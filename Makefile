default: test

depends:
	govendor sync

test: clean
	go test -cover -covermode=count -coverprofile=coverage.out ./...

test-silent: clean-build
	go test -cover -covermode=count -coverprofile=coverage.out ./... > dist/${CLI_VERSION}/unit.log

clean-build: clean
	rm -rf dist
	mkdir -p dist/${CLI_VERSION}/alpine

clean:
	rm -rf coverage.out

build: clean-build
	CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/alpine/apollo
	ls -la dist/${CLI_VERSION}/alpine/apollo

test-and-build: clean clean-build
	go test -cover -covermode=count -coverprofile=coverage.out ./...
	CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/alpine/apollo