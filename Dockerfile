FROM golang:latest

LABEL maintainer="Fabio Sebastiano <sebastiano.fabio@gmail.com"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8000

RUN go build

CMD ["./go-rest-api"]