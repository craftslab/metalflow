#!/bin/bash

TAG=13.1-alpine
PORT=5432

USER=postgres
PASS=postgres
DB=metalflow

docker pull postgres:$TAG
docker run --name postgres -p $PORT:5432 -e POSTGRES_USER=$USER -e POSTGRES_PASSWORD=$PASS -e POSTGRES_DB=$DB -d postgres:$TAG
