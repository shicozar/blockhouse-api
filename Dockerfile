# Step 1: Build the Go binary with cgo enabled
FROM golang:1.24-alpine AS builder

# Install build dependencies, including sqlite and gcc
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set the working directory in the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire project to the container
COPY . .

# Build the Go application with cgo enabled (no CGO_ENABLED=0)
RUN go build -o main .

# Step 2: Create the final image
FROM alpine:latest

# Install SQLite in the final image
RUN apk add --no-cache sqlite

# Set the working directory in the container
WORKDIR /root/

# Copy the built Go binary from the builder image
COPY --from=builder /app/main .

# Expose the port the app runs on
EXPOSE 8080

# Run the Go application
CMD ["./main"]
