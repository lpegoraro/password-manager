protoc --go_out=remote remote.proto
go build -a -o "$GOPATH"/bin/password-manager github.com/lpegoraro/password-manager/password-manager
