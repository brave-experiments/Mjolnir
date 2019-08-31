default: test

generate:
	go run builder/main.go

test: clean generate
	go test -cover -covermode=count -coverprofile=coverage.out ./...

test-silent: clean-build generate
	go test -cover -covermode=count -coverprofile=coverage.out ./... > dist/${CLI_VERSION}/unit.log

clean-build: clean
	rm -rf dist
	mkdir -p dist/${CLI_VERSION}/alpine

clean:
	rm -rf coverage.out

build: clean-build generate
	CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/alpine/apollo
	ls -la dist/${CLI_VERSION}/alpine/apollo

test-and-build: clean clean-build generate
	go test -cover -covermode=count -coverprofile=coverage.out ./...
	CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/alpine/apollo