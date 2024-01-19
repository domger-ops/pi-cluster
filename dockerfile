# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . .

# Build the Go app
RUN go build -o myapp

#Use a minimal base image for the final image
FROM --platform=linux/amd64,linux/arm64 alpine:latest

# Set the working directory to /app
WORKDIR /app

# Copy the binary from the builder image
COPY --from=builder /app/myapp .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]

