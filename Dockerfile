FROM golang:1.18.0-alpine3.15 as build

ENV CGO_ENABLED=0 \
    GOOS=linux

RUN apk add --no-cache curl

# Download and install go-swagger tool and swagger-ui.
RUN curl -L https://github.com/go-swagger/go-swagger/releases/download/v0.29.0/swagger_linux_amd64 -o /go/bin/swagger && \
    chmod a+x /go/bin/swagger && \
    curl -L https://github.com/swagger-api/swagger-ui/archive/refs/tags/v3.47.1.tar.gz | \
    tar -xz --strip-components 1 -C /go/bin swagger-ui-3.47.1/dist && \
    mv /go/bin/dist /go/bin/static && \
    sed -i 's,https://petstore.swagger.io/v2/swagger.json,swagger.yml,g' /go/bin/static/index.html

# Download generating tool.
RUN go install github.com/vektra/mockery/v2@v2.10.0

WORKDIR /go/src/github.com/AidarBabanov/billups-interview

# Download all project dependencies.
# This trick exploits how the docker uses the layers created on RUN statements.
# The dependencies will be re-downloaded only if the `go.mod` or `go.sum` was changed.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code, generate, test and build binaries.
COPY . .
RUN  go generate ./... && \
     go test ./... && \
    mkdir -p swagger && \
    swagger generate spec -o /go/bin/static/swagger.yml && \
    go install ./cmd/app

FROM alpine:3.15

WORKDIR /

# Copy swagger and application binary.
COPY --from=build /go/bin/app ./bin/app
COPY --from=build /go/bin/static ./static

CMD app