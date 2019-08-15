default: test

depends:
	govendor sync

test: clean
	go test -cover -covermode=count -coverprofile=coverage.out ./...

clean-build: clean
	rm -rf dist
	mkdir -p dist/${CLI_VERSION}/alpine

clean:
	rm -rf coverage.out

build: clean-build
	go build -a -installsuffix cgo -o dist/${CLI_VERSION}/alpine/apollo