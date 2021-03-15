FROM golang:latest

WORKDIR /app

ADD build/ /app

EXPOSE 7001-8001

ENTRYPOINT ["./password-manager"]