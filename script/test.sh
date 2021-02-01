#!/bin/bash

go env -w GOPROXY=https://goproxy.cn,direct

GIN_MODE=release go test -cover -v -parallel 2 ./...
