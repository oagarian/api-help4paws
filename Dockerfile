# Golang base image
FROM golang:latest

# Define workdir
WORKDIR /api/

# Copy the entire project
COPY . .

# Download dependencies
RUN go mod tidy

# Build the binary
RUN cd src && go build main.go

