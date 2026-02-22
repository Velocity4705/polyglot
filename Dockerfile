# Multi-stage build for Polyglot
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make

# Set working directory
WORKDIR /build

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o polyglot ./cmd/polyglot

# Final stage - include common language runtimes
FROM alpine:latest

# Install common language runtimes
RUN apk add --no-cache \
    python3 \
    nodejs \
    ruby \
    php \
    perl \
    lua5.3 \
    bash \
    gcc \
    g++ \
    musl-dev \
    go

# Create symlink for lua
RUN ln -s /usr/bin/lua5.3 /usr/bin/lua

# Copy binary from builder
COPY --from=builder /build/polyglot /usr/local/bin/polyglot

# Create working directory
WORKDIR /workspace

# Set entrypoint
ENTRYPOINT ["polyglot"]
CMD ["--help"]
