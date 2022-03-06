FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

ADD pkg ./pkg
ADD cmd ./cmd
ADD pkg /usr/local/go/src/pkg
RUN go install .
RUN go build -o /docker-gs-ping
EXPOSE 9010
# Install prerequisites
# RUN apk add --update \
  #  curl
CMD [ "/docker-gs-ping" ]