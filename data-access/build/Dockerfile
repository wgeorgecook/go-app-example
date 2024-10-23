# Dockerfile boilerplace generated using Gemini on gemini.google.com. 
# Release version 2024.10.09
# It was modified to pin specific base images and for specific 
# build directories.

# Build stage
FROM golang@sha256:db933bde327ab6b27f54c29b092ab3bf9276738432ed1bc730b5c9bd98ff33ef AS builder

WORKDIR /app

# Copy Go Modules and install dependencies 
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go binary statically 
RUN CGO_ENABLED=0 GOOS=linux go build -o data-access ./cmd/main.go

# Final stage
FROM alpine:3.20.0@sha256:77726ef6b57ddf65bb551896826ec38bc3e53f75cdde31354fbffb4f25238ebd

# Create apprunner user and group
RUN addgroup -S apprunner && adduser -S apprunner -G apprunner

WORKDIR /app

# Only copy the necessary binary and migrations files
COPY --from=builder /app/data-access /app/data-access
COPY --from=builder /app/internal/pkg/db/migrations /app/migrations

# Change ownership to apprunner user
RUN chown -R apprunner:apprunner /app

# Command to run the application as apprunner user
USER apprunner

CMD ["/app/data-access"]