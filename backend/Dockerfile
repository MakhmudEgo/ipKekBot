FROM golang:latest

WORKDIR /go/src/
COPY ./ ./

RUN go build main.go
CMD ["./main"]
EXPOSE 8080
