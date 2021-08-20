FROM golang:1.16 as build
WORKDIR /go/src/infinite-pong-manual-start/
COPY . .
RUN go build -o /build/infinite-pong-manual-start /go/src/infinite-pong-manual-start

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/infinite-pong-manual-start /

CMD ["/infinite-pong-manual-start"]