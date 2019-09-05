FROM iron/go:dev

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/lpegoraro/password-manager/
ADD . $SRC_DIR
RUN cd $SRC_DIR; go build -a -o $GOPATH/bin/password-manager github.com/lpegoraro/password-manager/password-manager

ENTRYPOINT ["./myapp"]