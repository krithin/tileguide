FROM docker.io/library/golang:latest AS builder

WORKDIR /go/src/app
COPY . .

RUN go install
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -o /go/bin/app ./...

##########################

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app

CMD ["/go/bin/app"]
