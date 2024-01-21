# Use the official Golang image
FROM golang:1.18

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOARCH=arm64 go build -o myapp

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./serverapp"]

