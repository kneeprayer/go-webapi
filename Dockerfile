FROM golang:1.9
EXPOSE 4000
WORKDIR /go/src/app
COPY server.go .
RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"] 