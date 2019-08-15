default: test

depends:
	govendor sync

test: clean
	go test -cover -covermode=count -coverprofile=coverage.out ./...

clean:
	rm -rf coverage.out