FROM golang:1.15

WORKDIR /go/src/dummy_app
RUN cd ..
COPY /dummy_app .

RUN go get -v ./...

RUN go install -v .

RUN ls /go/bin

ENTRYPOINT ["/go/bin/dummy_app"]