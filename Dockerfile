# Use Go 1.23 as the base image for building the Go app
FROM golang:1.23 AS build-env

# Copy the current directory contents to /app in the container
COPY . /app

# Set the working directory inside the container
WORKDIR /app

# Enable Go modules (this is the default behavior in Go 1.16+)
# Download dependencies and build the Go app
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o main .

# Use a minimal base image for the final container
FROM gcr.io/distroless/base

# Set the working directory in the final image
WORKDIR /app/

# Copy the binary from the build environment to the final image
COPY --from=build-env /app/main .

# Specify the command to run the binary
CMD ["./main"]
