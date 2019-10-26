ARG GO_VERSION=1.11

##########################################
FROM golang:${GO_VERSION}-alpine AS builder
WORKDIR /src

RUN apk add --no-cache git
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /my-app .

# Run tests
RUN go test -v ./...

# Production image stage
FROM scratch as FINAL

COPY --from=builder /my-app /my-app

ENTRYPOINT ["/my-app"]
