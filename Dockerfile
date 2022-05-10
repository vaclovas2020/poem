FROM golang:latest AS poem
WORKDIR /go
ADD . .
RUN go get && make
RUN rm -rf src
ENTRYPOINT ["/go/poem"]
