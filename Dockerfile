# Stage 1: build application
FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# the -C flag is used to specify the directory where the go build command should be executed from the root of the project
# the binary will be built in the cmd directory
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -C cmd/server -o balance

# Stage 2: production go image
FROM scratch

# default value if not provided during build
# it will be overridden by the value in the .env file (docker-compose)
ARG BALANCE_APP_PORT=8080
ENV BALANCE_APP_PORT=${BALANCE_APP_PORT}

WORKDIR /app

COPY --from=builder /app/cmd/server/balance .

EXPOSE ${BALANCE_APP_PORT}

ENTRYPOINT [ "./balance" ]