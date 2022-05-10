FROM golang:latest AS poem
WORKDIR /go
ADD . .
RUN cd src/app && go get && make
RUN rm -rf src
ENTRYPOINT ["/go/poem"]
