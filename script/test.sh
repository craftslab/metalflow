#!/bin/bash

go env -w GOPROXY=https://goproxy.cn,direct

go test -cover -v -parallel 2 ./...
