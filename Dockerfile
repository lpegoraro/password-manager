FROM iron/go:dev

WORKDIR /app

ENV SRC_DIR=/go/src/github.com/lpegoraro/password-manager/
ADD . $SRC_DIR
RUN cd $SRC_DIR; go build -o myapp; cd myapp /app/

ENTRYPOINT ["./myapp"]