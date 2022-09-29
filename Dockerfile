#syntax=docker/dockerfile:1.2

###############################################################################
# Golang builder 

FROM golang:1.19.0-buster AS go-builder
WORKDIR /app

# Pre-copy/cache go.mod for pre-downloading dependencies and only 
# redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code to builder
COPY . .

RUN go test ./...

###############################################################################
# Build base for app

FROM go-builder as app-build

RUN CGO_ENABLED=0 GOOS=linux go build -v -a -o /usr/local/app ./cmd/main.go

###############################################################################
# Build final app image

FROM alpine:3.12.4 AS app
WORKDIR /usr/local/bin
COPY --from=app-build /usr/local/app .
RUN chmod +x ./app

CMD ["app"] 

###############################################################################
# Build final test app image

FROM go-builder as test-app
WORKDIR /app/hack
ENTRYPOINT ["./e2e_tests.sh"]

