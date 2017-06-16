FROM golang:1.8

WORKDIR /go/src/calculator
COPY . .

RUN go build -buildmode=plugin -o plugins/add.so add.go && \
    go build -buildmode=plugin -o plugins/mult.so mult.go

RUN go build

CMD ["./calculator"]
