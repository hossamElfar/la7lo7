FROM golang:latest

RUN mkdir /usr/local/go/src/la7lo7

ADD . /usr/local/go/src/la7lo7

WORKDIR /usr/local/go/src/la7lo7

RUN go build -o main .

CMD ["/usr/local/go/src/la7lo7/main"]