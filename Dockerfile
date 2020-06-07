FROM golang:1.14.2 AS build-env

ENV GO111MODULE=on

ADD . /apiserver

WORKDIR /apiserver

RUN ls -lrt

RUN go mod download

COPY . .

RUN go build server.go

EXPOSE 8989

CMD ["./server"]