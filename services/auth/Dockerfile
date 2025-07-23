# syntax=docker/dockerfile:1

FROM golang:1.24-alpine

WORKDIR /app

# Copy go mod & download dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy entire source
COPY . .

# Build binary
RUN go build -o main .

# Run binary
CMD ["./main"]
