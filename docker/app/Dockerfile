# Development environment
# -----------------------
FROM golang:1.18-alpine
WORKDIR /myapp

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app \
    && go build -ldflags '-w -s' -a -o ./bin/migrate ./cmd/migrate \
    && chmod +x /myapp/docker/app/bin/*

EXPOSE 8080