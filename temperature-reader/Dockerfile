FROM golang:1.16 as build
WORKDIR /go/src/infinite-pong/
COPY . .
RUN go build -o /build/infinite-pong /go/src/infinite-pong

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/infinite-pong /

CMD ["/infinite-pong"]