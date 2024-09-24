# Dockerfile

FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0
RUN go build -o /server ./cmd

# Use uma imagem base diferente
FROM ubuntu:22.04

# Instale glibc
RUN apt-get update && apt-get install -y libc6

COPY --from=builder /server /server

CMD ["/server"]
