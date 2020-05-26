FROM golang:latest
RUN go get github.com/labstack/echo/...
RUN mkdir /go/src/work
WORKDIR /go/src/work
ADD . /go/src/work
CMD ["go", "run", "main.go"]
