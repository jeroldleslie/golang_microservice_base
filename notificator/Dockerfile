FROM golang

RUN mkdir -p /go/src/go-microservice-base

ADD . /go/src/go-microservice-base

RUN curl https://glide.sh/get | sh
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN cd /go/src/go-microservice-base && glide install

ENTRYPOINT  watcher -run go-microservice-base/notificator/cmd -watch go-microservice-base/notificator
