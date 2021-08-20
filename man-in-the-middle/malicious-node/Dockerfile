FROM golang:1.16 as build
WORKDIR /go/src/malicious-node/
COPY .. .
RUN go build -o /build/malicious-node /go/src/malicious-node

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/malicious-node /

CMD ["/malicious-node"]