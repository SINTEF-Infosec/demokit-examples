FROM golang:1.16 as build
WORKDIR /go/src/pong-with-limit/
COPY . .
RUN go build -o /build/pong-with-limit /go/src/pong-with-limit

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/pong-with-limit /

CMD ["/pong-with-limit"]