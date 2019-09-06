FROM golang:latest

WORKDIR /app

ADD password-manager /app/

ENTRYPOINT ["./password-manager"]