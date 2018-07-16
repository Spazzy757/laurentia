FROM golang:latest as build

WORKDIR $GOPATH/src/app
COPY . .

RUN go version && go get -u -v golang.org/x/vgo
RUN vgo build ./...

FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
CMD ["/app"]