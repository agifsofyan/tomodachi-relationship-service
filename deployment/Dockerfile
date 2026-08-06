# ---------- Builder ----------
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache \
    ca-certificates \
    tzdata

WORKDIR /app

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /app/bin/relationship-service ./cmd/server

# ---------- Runtime ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates && \
    addgroup -S appgroup && \
    adduser -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/bin/relationship-service .

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["./relationship-service"]