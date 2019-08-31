# Build Apollo in a stock Go builder container
FROM golang:1.12-alpine3.10 as apollo-builder

RUN apk add --no-cache make \
    gcc \
    musl-dev \
    linux-headers\
    git \
    curl

RUN go get github.com/githubnemo/CompileDaemon

VOLUME /usr/local/go/src/github.com/brave-experiments/apollo-devops
WORKDIR /usr/local/go/src/github.com/brave-experiments/apollo-devops

ADD . .
RUN go get -v
RUN apk add openssh
RUN ssh-keygen -t rsa -N "" -f ~/.ssh/id_rsa
RUN make generate

CMD CompileDaemon -log-prefix=false -build="make build" -command="./apollo"
