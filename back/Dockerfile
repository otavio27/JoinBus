FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . /app
RUN go build -o /JoinBus ./cmd/main.go
CMD [ "/JoinBus" ]