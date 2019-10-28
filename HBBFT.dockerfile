FROM parity/parity:stable as notifier

FROM rust:slim as builder

LABEL maintainer="sdare@brave.com"


RUN apt-get update -qq && apt-get install build-essential cmake git libudev-dev -qqy && \
    git clone https://github.com/poanetwork/parity-ethereum.git --branch hbbft  && \
    cd parity-ethereum && \
    cargo build --release --features final


FROM ubuntu:latest

COPY --from=builder /parity-ethereum/target/release/parity /bin/

ENTRYPOINT ["/bin/parity"]