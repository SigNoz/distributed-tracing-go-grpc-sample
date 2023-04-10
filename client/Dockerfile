# use a minimal alpine image
FROM golang:alpine3.17 AS builder

# Buildkit arguments for cross-compilation
ARG TARGETOS=linux TARGETARCH=amd64

# set working directory and copy source code
WORKDIR /go/src/github.com/vabira200/go-grpc-crud-example

# Cache dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download -x

# Add the sources and proceed with build
COPY . .

# Build the binary
RUN go build -o ./bin/client ./client


# Use a minimal alpine image
FROM alpine:3.17

# Add Maintainer Info
LABEL maintainer="signoz"

# Add ca-certificates in case you need to call HTTPS endpoints
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

# Copy the binary from builder image
COPY --from=builder /go/src/github.com/vabira200/go-grpc-crud-example/bin/client .

# run the binary as the entrypoint
ENTRYPOINT [ "/client" ]