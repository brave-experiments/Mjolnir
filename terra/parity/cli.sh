#!/usr/bin/env bash

cp .env.dist .env
docker-compose up -d
docker-compose run cli bash