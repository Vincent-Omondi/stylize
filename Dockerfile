# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Metadata
LABEL maintainer="oumaphilip01@gmail.com"
LABEL version="1.0"
LABEL description="A web-based application built with Go, CSS, HTML, and JavaScript"

# Set the Current Working Directory inside the container
WORKDIR /app


# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# Install git and other dependencies (if needed)
RUN apk update && apk add --no-cache git
RUN apk update && apk add --no-cache bash
RUN go mod init stylize


# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
