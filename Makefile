# Builds ImmuDB as docker and Install password manager locally in the GOPATH bin directory

.PHONY: all
all: docker-immudb install

docker-immudb-teardown:
	$(info =================== Cleaning ImmuDB Docker Container ===================)
	docker stop immudb
	docker remove immudb
	docker network remove immudb

docker-immudb:
	$(info =================== Starting ImmuDB Docker Container ===================)
	docker network create immudbnet
	docker run -d --net immudbnet -it --rm --name immudb -p 3322:3322 codenotary/immudb:latest

install:
	$(info =================== Installing Password Manager ===================)
	protoc --go-grpc_out=remote remote.proto
	go build -o "${GOPATH}"/bin/password-manager github.com/lpegoraro/password-manager/password-manager

install-protoc:
	$(info =================== Installing Protoc ===================)
	#go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    #go get -u google.golang.org/grpc)
