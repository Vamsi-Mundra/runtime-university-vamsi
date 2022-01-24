FROM golang:1.14-alpine

ENV SERVER_URL=grpc-server.herokuapp.com:80

WORKDIR /go/src/app
COPY . .

RUN go install ./...

CMD ["cmd", "-lat", "407838351", "-long", "-746143763"]