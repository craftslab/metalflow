FROM golang:latest AS build-stage
WORKDIR /go/src/app
COPY . .
RUN apt update && \
    apt install -y upx
RUN make build

FROM gcr.io/distroless/base-debian10 AS production-stage
WORKDIR /
COPY --from=build-stage /go/src/app/bin/* /
COPY --from=build-stage /go/src/app/config/*.yml /
USER nonroot:nonroot
