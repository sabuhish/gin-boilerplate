FROM golang:latest


ENV GIN_MODE=release

RUN go build .


# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN go build .

# Expose port 8080 to the outside world
EXPOSE 8000

#Command to run the executable
CMD ["./main"]