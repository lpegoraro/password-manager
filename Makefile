# Builds ImmuDB as docker and Install password manager locally in the GOPATH bin directory

.PHONY: all
all: docker-immudb install

docker-immudb-teardown:
	$(info =================== Cleaning ImmuDB Docker Container ===================)
	docker rm -f immudb || echo "No container found"
	docker network prune -a -y

docker-immudb:
	$(info =================== Starting ImmuDB Docker Container ===================)
	docker network create immudbnet || echo "Network already created"
	docker run -d --net immudbnet -it --rm --name immudb -p 3322:3322 codenotary/immudb:latest || echo " Container was up"

install:
	$(info =================== Installing Password Manager ===================)
	# protoc --go-grpc_out=. remote.proto
	go build -o "${GOPATH}"/bin/password-manager github.com/lpegoraro/password-manager/password-manager

install-protoc:
	$(info =================== Installing Protoc ===================)
	#go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    #go get -u google.golang.org/grpc)
