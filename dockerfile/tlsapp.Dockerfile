FROM golang:1.15

WORKDIR /go/src/metrictlsapp
RUN cd ..
COPY /metrictlsapp .

COPY ../cert cert/

RUN go get -v ./...

RUN go install -v .

RUN ls /go/bin

ENTRYPOINT ["/go/bin/metrictlsapp"]