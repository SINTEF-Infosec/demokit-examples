FROM golang:1.16 as build
WORKDIR /go/src/temperature-reader/
COPY . .
RUN go build -o /build/temperature-reader /go/src/temperature-reader

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/infinite-pong /

CMD ["/temperature-reader"]