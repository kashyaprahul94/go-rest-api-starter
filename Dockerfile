##################
# Builder region #
##################

# Get golang image to build the project
FROM golang:1.13-stretch as builder

# Perform all operations inside `/app` directory
WORKDIR /app

# Copy all contents from source directory to work directory
COPY . /app

# Build the project to generate binary for `linux` architecture
RUN CGO_ENABLED=0 GOOS=linux make build


##################
# Runner region #
##################

# To run the binary, use alpine image ( super tiny )
FROM alpine:latest

# Perform all operations inside `/app` directory
WORKDIR /app

# Expose the http server port environment variable
ENV PORT 4111

# Copy the binary from build stage to `/app` directory
COPY --from=builder /app/bin .

# Invoke the binary
CMD ["./rest-api"]