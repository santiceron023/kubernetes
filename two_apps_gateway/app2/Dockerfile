# Dockerfile
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Only copy go.mod if go.sum is not present
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o app2 .

# Final image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app2 .

CMD ["./app2"]
