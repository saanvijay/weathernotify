# Use an Ubuntu base image
FROM ubuntu:22.04

# Install dependencies
RUN apt-get update && \
    apt-get install -y wget software-properties-common apt-transport-https ca-certificates gcc g++ libc-dev cmake libssl-dev zlib1g-dev

# Build and install librdkafka from source
RUN wget https://github.com/edenhill/librdkafka/archive/refs/tags/v1.9.2.tar.gz && \
    tar -xvzf v1.9.2.tar.gz && \
    cd librdkafka-1.9.2 && \
    ./configure --prefix=/usr && make && make install && \
    cd .. && rm -rf librdkafka-1.9.2 v1.9.2.tar.gz

# Download and install Go 1.23
RUN wget https://go.dev/dl/go1.23.0.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.23.0.linux-amd64.tar.gz && \
    rm go1.23.0.linux-amd64.tar.gz

# Set up Go environment variables
ENV PATH="/usr/local/go/bin:${PATH}"

# Set up the working directory
WORKDIR /app

# Copy go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application with CGO
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o weatherproducer .

EXPOSE 8080

# Run the producer binary by default
ENTRYPOINT ["/app/weatherproducer"]
