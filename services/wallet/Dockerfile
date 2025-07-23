# syntax=docker/dockerfile:1

FROM golang:1.24-alpine

# Install git dan openssl untuk modul & TLS
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files dan download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o main .

# Expose port sesuai .env
EXPOSE 8082

# Run binary
CMD ["./main"]
