FROM golang:latest

WORKDIR /app

ADD password-manager /app/
EXPOSE 7001-8001

ENTRYPOINT ["./password-manager"]