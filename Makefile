default: docker-test

static-switch:
	cp terra/static.go.dist terra/static.go

generate: static-switch
	go run builder/main.go

docker-test: clean generate
	go test -cover -covermode=count -coverprofile=coverage.out ./...

docker-test-silent: clean-build generate
	go test -cover -covermode=count -coverprofile=coverage.out ./... > dist/${CLI_VERSION}/unit.log

test-silent-connection: clean-build
	go test -cover -covermode=count -coverprofile=coverage.out ./connect/... > dist/${CLI_VERSION}/unit.log

clean-build: clean
	rm -rf dist/${CLI_VERSION}/unix
	mkdir -p dist/${CLI_VERSION}/unix

clean-build-mac: clean
	rm -rf dist/${CLI_VERSION}/osx
	mkdir -p dist/${CLI_VERSION}/osx

clean:
	rm -rf coverage.out

test-and-build: clean clean-build generate
	go test -cover -covermode=count -coverprofile=coverage.out ./...
	GOPROXY=https://proxy.golang.org CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/unix/mjolnir

restart:
	docker-compose down --remove-orphans
	docker-compose up -d

copy:
	cp docker-compose.override.yml.dist docker-compose.override.yml
	cp .env.dist .env

dev: copy restart
	docker-compose exec cli sh

TARGET := $(shell uname)

create: create-$(TARGET)

create-Darwin: generate clean-build-mac
	# GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/osx/mjolnir
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o mjolnir
	# cp dist/${CLI_VERSION}/osx/mjolnir . 
	# ls -la dist/${CLI_VERSION}/osx/mjolnir

create-linux: generate clean-build
	# CGO_ENABLED=0 go build -a -installsuffix cgo -o dist/${CLI_VERSION}/unix/mjolnir
	CGO_ENABLED=0 go build -a -installsuffix cgo -o mjolnir
	# cp dist/${CLI_VERSION}/unix/mjolnir . 
	# ls -la dist/${CLI_VERSION}/unix/mjolnir

build: copy restart 
	docker-compose exec -T cli make create TARGET=$(TARGET)

quorum: 
	./mjolnir apply quorum examples/values-local.yml

destroy:
	./mjolnir destroy examples/values-local.yml

test-ci: 
	cp docker-compose.override.test.yml.dist docker-compose.override.yml
	docker-compose up -d --no-deps cli-test
	sleep 2
	docker-compose exec -T cli-test make docker-test

tests-watch:
	cp docker-compose.override.test.yml.dist docker-compose.override.yml
	docker-compose up --no-deps cli-test

tests-silent:
	cp docker-compose.override.test.yml.dist docker-compose.override.yml
	docker-compose up -d --no-deps cli-test
	sleep 2
	docker-compose exec -T cli-test make docker-test-silent