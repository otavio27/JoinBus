FROM golang

WORKDIR /app
COPY . /app
RUN go build -o /JoinBus ./cmd/main.go
CMD [ "/JoinBus" ]