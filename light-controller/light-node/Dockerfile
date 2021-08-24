FROM golang:1.16 as build
WORKDIR /go/src/light-node/
COPY . .
RUN go build -o /build/light-node /go/src/light-node

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/light-node /

CMD ["/light-node"]