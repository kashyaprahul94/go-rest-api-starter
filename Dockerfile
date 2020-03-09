##################
# Builder region #
##################

# Get golang image to build the project
FROM golang:1.13-stretch as builder

# Perform all operations inside `/app` directory
WORKDIR /app

# Copy only dependencies related files first
COPY go.mod .
COPY go.sum .
COPY Makefile .

# Download all deps first
RUN make install

# Copy all contents from source directory to work directory
COPY . .

# Build the project to generate binary for `linux 64` architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 make build


##################
# Runner region #
##################

# To run the binary, use `scratch` image ( super tiny )
FROM scratch

# Perform all operations inside `/app` directory
WORKDIR /app

# Expose the http server port environment variable
ENV PORT 4111

# Copy the binary from build stage to `/app` directory
COPY --from=builder /app/bin .

# Invoke the binary
CMD ["./rest-api"]