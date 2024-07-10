# Start with a base image that includes Golang tools necessary to compile Go code
FROM golang:1.23-rc as builder

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy the source code to the working directory
COPY . .

# Build the Go application with specific build flags
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o data_ingest_worker ./cmd/main.go

# Start a new stage from a minimal base image
FROM scratch

# Copy the built binary from the previous stage into the container
COPY --from=builder /usr/src/app/data_ingest_worker ./data_ingest_worker

# Copy the data.json file from the build stage to the root directory of the container
COPY ./adapters/db/data.json ./data.json

EXPOSE 4001
# Set the command to run when the container starts
CMD ["./data_ingest_worker"]


