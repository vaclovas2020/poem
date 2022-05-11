FROM golang:latest AS poem
WORKDIR /go
ADD . ./src/app
RUN cd src/app && go get && go build -o /go/poem /go/src/app/cmd/main/main.go
RUN rm -rf src
ENTRYPOINT ["/go/poem"]
EXPOSE 7789