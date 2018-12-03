FROM golang

RUN mkdir -p /go/src/fivekilometer

ADD . /go/src/fivekilometer

RUN curl https://glide.sh/get | sh
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

RUN cd /go/src/fivekilometer && glide install

ENTRYPOINT  watcher -run fivekilometer/notificator/cmd -watch fivekilometer/notificator
