#!/bin/bash

VERSION=3.4.14
ETCD=etcd-v$VERSION-linux-amd64
URL=http://127.0.0.1:2379

curl -L https://github.com/etcd-io/etcd/releases/download/v$VERSION/$ETCD.tar.gz -o $ETCD.tar.gz
tar zxvf $ETCD.tar.gz

export PATH=$PWD/$ETCD:$PATH

# Start a local etcd server
etcd --advertise-client-urls $URL --listen-client-urls $URL &

# Write to etcd
#ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 put "/path/to/key" "val"

# Read from etcd
#ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 get "/" --keys-only=true --prefix
#ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 get "/path/to/" --prefix
#ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 get "/path/to/key"

# Delete from etcd
#ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 del "/path/to/key"

# Watch on etcd
#ETCDCTL_API=3 etcdctl --endpoints=127.0.0.1:2379 watch "/path/to" --prefix
