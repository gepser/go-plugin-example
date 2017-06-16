FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go build -buildmode=plugin -o plugins/mysql.so mysql.go && \
    ls
    #go build -buildmode=plugin -o plugins/mult.so mult.go

RUN go build

CMD ["./app"]
