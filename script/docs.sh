#!/bin/bash

# USAGE: https://github.com/swaggo/swag/blob/master/README.md
# WEB: http://127.0.0.1:9080/swagger/index.html

release=1.7.0

if [ -d swag ]; then
    rm -rf swag
fi

mkdir swag

curl -L https://github.com/swaggo/swag/releases/download/v${release}/swag_${release}_Linux_x86_64.tar.gz -o swag.tar.gz
tar zxvf swag.tar.gz -C swag/
rm -rf swag.tar.gz

./swag/swag init

rm -rf swag
