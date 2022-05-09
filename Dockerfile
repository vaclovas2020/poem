FROM golang:latest AS poem
WORKDIR /go
ADD . ./src/app
RUN cd src/app && go get && make
RUN rm -rf src
ENTRYPOINT ["bin/poem"]
