FROM golang:1.20

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y librdkafka-dev

RUN go mod download -x
RUN go build -o walletcore ./cmd/walletcore/main.go
# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C cmd/walletcore -o walletcore

CMD ["./walletcore"]