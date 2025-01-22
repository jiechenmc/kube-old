############################
# STEP 1 build executable binary
############################
FROM golang:1.23-alpine AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -o kube

############################
# STEP 2 build a small image
############################
FROM alpine:latest

COPY --from=builder /app/kube /app/kube

ENV PORT=8080
ENV GIN_MODE=release
EXPOSE 8080

CMD ["/app/kube"]