#!/bin/bash
docker build -t go-webserver .
docker run --publish 8080:8080 --name go-test --rm go-webserver 