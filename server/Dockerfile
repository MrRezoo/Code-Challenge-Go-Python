# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

LABEL name="GoServer"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# Stage 2: Run the Go application
FROM alpine:latest
LABEL name="GoServer"

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/config ./config

CMD ["./main"]
