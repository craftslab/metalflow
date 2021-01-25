# Config

VERSION=$(version)


# Build

.PHONY: FORCE

build: go-build
.PHONY: build

clean: go-clean
.PHONY: clean

docs: go-docs
.PHONY: docs

etcd: go-etcd
.PHONY: etcd

lint: go-lint
.PHONY: lint

proto: go-proto
.PHONY: proto

test: go-test
.PHONY: test


# Non-PHONY targets (real files)

go-build: FORCE
	./script/build.sh $(VERSION)

go-clean: FORCE
	./script/clean.sh

go-docs: FORCE
	./script/docs.sh

go-etcd: FORCE
	./script/etcd.sh

go-lint: FORCE
	./script/lint.sh

go-proto: FORCE
	./script/proto.sh

go-test: FORCE
	./script/test.sh
