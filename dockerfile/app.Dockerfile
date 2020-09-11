FROM golang:1.15

WORKDIR /go/src/metricapp
RUN cd ..
COPY /metricapp .

RUN go get -v ./...

RUN go install -v .

RUN ls /go/bin

ENTRYPOINT ["/go/bin/metricapp"]