FROM golang:1.8

WORKDIR /go/src
COPY  . github.com/benwaine/artiste

WORKDIR /go/src/github.com/benwaine/artiste/srv/artiste

RUN go-wrapper install

EXPOSE 8082

ENTRYPOINT /go/bin/artiste --config=/go/src/github.com/benwaine/artiste/srv/artiste/test/test_conf.json