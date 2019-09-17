protoc --go_out=plugins=grpc:remote remote.proto
go build -a -o "$GOPATH"/bin/password-manager github.com/lpegoraro/password-manager/password-manager
