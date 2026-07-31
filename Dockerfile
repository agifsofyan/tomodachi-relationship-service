# --- Contoh Dockerfile untuk service Golang (relationship-service) ---
# Letakkan file ini di root folder relationship-service/ (sejajar dengan go.mod)
# Multi-stage build: hasil akhir image kecil, tanpa toolchain Go.

# Stage 1: build binary
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Sesuaikan path "./cmd/server" dengan lokasi package main kamu
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/relationship-service ./cmd/server

# Stage 2: runtime image minimal
FROM alpine:3.20

RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/bin/relationship-service .

EXPOSE 8080

CMD ["./relationship-service"]
