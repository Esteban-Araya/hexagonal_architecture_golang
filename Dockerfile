# Use the official Golang image as the base image
FROM golang:1.22.3-bookworm

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifest
COPY . .

# Build the Go app
# RUN go build -o websocket-server

# Install air for live reloading
RUN go install github.com/air-verse/air@v1.52.3 \
    && go install go.uber.org/mock/mockgen@v0.4.0


RUN go mod tidy
# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["air"]