FROM golang:1.7.1-alpine as build

WORKDIR /go/src

COPY . .

RUN GOOS=linux go build main.go

CMD [ "./main" ]