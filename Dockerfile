FROM golang:latest

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

RUN go build -o main code.go

CMD ["./main"]

