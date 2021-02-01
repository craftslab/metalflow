FROM golang:latest as build-stage
WORKDIR /go/src/app
COPY . .
RUN apt update && \
    apt install -y upx
RUN make build

FROM nginx as production-stage
WORKDIR /go/dist/
RUN mkdir -p /go/dist/bin && \
    mkdir -p /go/dist/etc
COPY --from=build-stage /go/src/app/bin/* /go/dist/bin/
COPY --from=build-stage /go/src/app/config/*.yml /go/dist/etc/
