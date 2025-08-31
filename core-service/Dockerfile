FROM golang:1.20

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y librdkafka-dev

RUN go mod download -x
RUN go build -o walletcore ./cmd/walletcore/main.go

CMD ["./walletcore"]
