# Stage 1: Build stage
FROM golang:alpine AS builder

# Install git
RUN apk update && apk add --no-cache git

# Set the working directory
WORKDIR /src

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o trx-core-svc main.go

# Stage 2: Final stage
FROM alpine:3.19

# Set the working directory
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /src/trx-core-svc .

# Expose port 5000
EXPOSE 5000

# Command to run the executable
CMD ["./trx-core-svc"]
