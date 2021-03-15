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
	docker run -it -d -p 3322:3322 -p 9497:9497 -v immudb:/var/lib/immudb --env IMMUDB_ADDRESS=0.0.0.0 --name immudb codenotary/immudb:latest || echo " Container was up"

proto-gen:
	$(info =================== Starting ImmuDB Docker Container ===================)
	# For unknown reasons it is not longer generating messages as AddPasswordReq, etc...
	# protoc --go-grpc_out=. -o=remote remote.proto

install:
	$(info =================== Installing Password Manager ===================)
	go build -o build/password-manager github.com/lpegoraro/password-manager/password-manager
	cp build/password-manager "${GOPATH}"/bin/password-manager

install-protoc:
	$(info =================== Installing Protoc ===================)
	#go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
    #go get -u google.golang.org/grpc)

build-certificate:
ifeq (,$(wildcard certs))
	$(info =================== Generate certificate for Signing Passwords ===================)
	mkdir certs
	openssl genrsa -out certs/pwdmgr.key 2048
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout certs/pwdmgr.key -out certs/pwdmgr.crt  -subj "/C=US/ST=Georgia/L=Atlanta/O=PasswordManager/CN=localhost"
endif
