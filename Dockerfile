# Use the official Golang image to create a build artifact.
# This is the first stage, called "builder".
FROM golang:1.16 AS builder

# Create and change to the app directory.
WORKDIR /app

# Copy the Go modules manifests.
COPY go.mod go.sum ./

# Download the dependencies.
RUN go mod download

# Copy the source code.
COPY . .

# Build the application.
RUN go build -o server

# Use the official Debian image for a lean production stage.
FROM debian:buster

# Copy the build artifact from the builder stage.
COPY --from=builder /app/server /server

# Expose port 80 to the outside world
EXPOSE 80

# Run the web service on container startup.
CMD ["/server"]