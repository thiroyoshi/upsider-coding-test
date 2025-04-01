# syntax=docker/dockerfile:1

# Use the official Golang image as Development image
FROM golang:1.23.1-bullseye AS development

# Set destination for COPY
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY go.mod go.sum ./
COPY ./cmd ./cmd
COPY ./internal ./internal

# Download dependencies
RUN go mod download && \
    go install github.com/air-verse/air@latest && \
    go install github.com/go-delve/delve/cmd/dlv@latest

WORKDIR /app/cmd/api

CMD ["air", "-c", ".air.toml"]

# Use the official Golang image as the base image
FROM golang:1.23.1-bullseye AS builder

# Set destination for COPY
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY go.mod go.sum ./
COPY ./cmd ./cmd
COPY ./internal ./internal

# Download dependencies
RUN go mod download

# Build the Go app
WORKDIR /app/cmd/api
RUN go build -o /dist/api

# Start a new stage from scratch
FROM gcr.io/distroless/base-debian12:nonroot AS production

ARG ENV

WORKDIR /app

# Expose port on which the app runs
EXPOSE 8080

# Copy the binary from builder
COPY --from=builder /dist/api /app/api

# Copy the viper config file
COPY ./cmd/api/viper/${ENV}.yaml ./viper/${ENV}.yaml

# Run the binary
CMD ["/app/api"]
