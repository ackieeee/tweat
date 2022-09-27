FROM golang:1.18.0 as builder

RUN apt update && apt upgrade -y && mkdir -p /var/log/ap
RUN go install github.com/cosmtrek/air@v1.29.0 && go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app

CMD ["air", "-c", ".air.toml"]
