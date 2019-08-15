# Build Apollo in a stock Go builder container
FROM golang:1.12-alpine3.10 as apollo-builder

RUN apk add --no-cache make \
    gcc \
    musl-dev \
    linux-headers\
    git \
    curl

RUN go get github.com/githubnemo/CompileDaemon \
    github.com/kardianos/govendor

VOLUME /usr/local/go/src/github.com/brave-experiments/apollo-devops
WORKDIR /usr/local/go/src/github.com/brave-experiments/apollo-devops

ADD . .
RUN govendor sync

CMD CompileDaemon -log-prefix=false -build="go build -a -installsuffix cgo -o apollo" -command="./apollo"
