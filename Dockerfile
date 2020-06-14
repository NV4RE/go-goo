############################
# STEP 1 setup for build project
############################
FROM golang:1.14.2 AS builder
WORKDIR $GOPATH/src/github.com/nv4re/app/
COPY . .
# Enable go module
ENV GO111MODULE=on
# Fetch dependencies
RUN go mod download
# Build the binary
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/app

############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
COPY --from=builder /go/bin/app /go/bin/app
# Run the gogo-blueprint binary
ENTRYPOINT ["/go/bin/app"]
