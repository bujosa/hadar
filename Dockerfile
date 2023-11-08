# Author: David Bujosa
# https://hub.docker.com/layers/library/golang/1.21.3-alpine
FROM golang:1.21.3-alpine 

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./.

RUN go build -o /myapp

EXPOSE 8080

CMD [ "/myapp" ]