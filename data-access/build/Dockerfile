# Dockerfile boilerplace generated using Gemini on gemini.google.com. 
# Release version 2024.10.09
# It was modified to pin specific base images and for specific 
# build directories.

# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY . .

RUN go build -o data-access ./cmd/main.go

# Run stage
FROM alpine@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d

WORKDIR /app

COPY --from=builder /app/data-access .

EXPOSE 8088

CMD ["/app/data-access"]