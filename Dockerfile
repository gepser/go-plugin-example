FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN "go build -buildmode=plugin -o add.so add.go"
RUN "go build"

CMD ["go-plugin-calculator"]
