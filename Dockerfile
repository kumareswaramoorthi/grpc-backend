FROM golang:1.15

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get ./...

# Install the package
RUN go install ./...

RUN go build .

# This container exposes port 8080 to the outside world
EXPOSE 8081

# Run the executable
CMD ["./grpc-envoy"]