FROM golang:latest

COPY code.go /go/src/app/code.go

COPY input.txt /go/src/app/input.txt

RUN mkdir -p /go/src/app

WORKDIR /go/src/app

RUN go build -o main code.go

CMD ["./main"]

