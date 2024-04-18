#!/bin/bash
set -e

trap 'killall main' SIGINT

cd $(dirname $0)

killall main || true
sleep 0.1


go run main.go -db-location=moscow.db -http-addr=127.0.0.1:8080 -config-file=sharding.toml -shard=Moscow &
go run main.go -db-location=moscow-r.db -http-addr=127.0.0.1:8081 -config-file=sharding.toml -shard=Moscow -replica &

go run main.go -db-location=minsk.db -http-addr=127.0.0.1:8082 -config-file=sharding.toml -shard=Minsk &
go run main.go -db-location=minsk-r.db -http-addr=127.0.0.1:8083 -config-file=sharding.toml -shard=Minsk -replica &

go run main.go -db-location=kiev.db -http-addr=127.0.0.1:8084 -config-file=sharding.toml -shard=Kiev &
go run main.go -db-location=kiev-r.db -http-addr=127.0.0.1:8085 -config-file=sharding.toml -shard=Kiev -replica &

go run main.go -db-location=tashkent.db -http-addr=127.0.0.1:8086 -config-file=sharding.toml -shard=Tashkent &
go run main.go -db-location=tashkent-r.db -http-addr=127.0.0.1:8087 -config-file=sharding.toml -shard=Tashkent -replica &

wait