.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/trinityfx/trinity/cmd/swarm
	go build -o bin/geth     github.com/trinityfx/trinity/cmd/geth
	go build -o bin/bootnode github.com/trinityfx/trinity/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/trinityfx/trinity/swarm/...
