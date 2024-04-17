# Start from the official Go Lang image
FROM golang:1.22.2-bullseye

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go Lang source code into the container
COPY . .

# Build the Go Lang application
RUN go mod download && go build -o zocket-key-value-store ./cmd/server

# Expose the port the application runs on
EXPOSE 8056

# Command to run the executable
CMD ["./zocket-key-value-store"]
