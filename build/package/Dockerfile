FROM golang:1.19.1-alpine3.16 AS src
RUN apk update && apk add --no-cache git
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .

FROM golangci/golangci-lint:v1.49.0-alpine AS lint
WORKDIR /workspace
COPY --from=src /workspace ./
RUN golangci-lint run

FROM src AS test
RUN CGO_ENABLED=0 go test ./... -coverprofile cover.out && go tool cover -func cover.out

FROM src AS build-server
RUN cd cmd/server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o server

FROM src AS build-client
RUN cd cmd/client && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o client

FROM alpine:3.16.2 AS release-base
ARG CACHEBUST
RUN apk update \
    && apk add --no-cache \
        ca-certificates \
        tzdata \
    && update-ca-certificates
RUN adduser -D -u 10001 appuser
USER appuser
WORKDIR /app

FROM release-base AS release-server
ENTRYPOINT ["./server"]
COPY --from=build-server /workspace/cmd/server/server ./server

FROM release-base AS release-client
ENTRYPOINT ["./client"]
CMD ["--help"]
COPY --from=build-client /workspace/cmd/client/client ./client