FROM golang:1.25.3-alpine3.22

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -ldflags="-s -w" -o binary

ENTRYPOINT [ "/app/binary" ]