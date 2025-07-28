FROM golang:1.24.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o bin/app cmd/main.go

FROM alpine:latest

WORKDIR /app

RUN mkdir etc
COPY --from=builder /app/bin/app .
COPY --from=builder /app/etc/config.yml etc/config.yml


EXPOSE 8080
ENTRYPOINT ["./app", "-config=/etc"]

