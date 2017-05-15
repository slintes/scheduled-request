#!/usr/bin/env bash

echo "building linux binary"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o scheduled-request-linux .

