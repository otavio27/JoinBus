FROM golang:1.20-alpine as builder

COPY back /app

WORKDIR /app

RUN go build -o /JoinBus ./cmd/main.go

FROM alpine:3.18

COPY --from=builder /JoinBus /JoinBus

CMD ["/JoinBus"]