# Second stage: Build the Go application
FROM golang:1.23-rc as builder

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app


# Copy the source code to the working directory
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o data_store_service ./cmd/main.go

# Start a new stage from a minimal base image
FROM scratch

# Copy the built binary from the previous stage into the container
COPY --from=builder /usr/src/app/data_store_service ./data_store_service

EXPOSE 8080
# Set the command to run when the container starts
CMD ["./data_store_service"]

