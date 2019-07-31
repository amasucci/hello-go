
FROM golang

ADD . /go/src/github.com/amasucci/hello-go

RUN go install github.com/amasucci/hello-go

ENTRYPOINT /go/bin/hello-go
