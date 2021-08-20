FROM golang:1.16 as build
WORKDIR /go/src/legit-node/
COPY .. .
RUN go build -o /build/legit-node /go/src/legit-node

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/legit-node /

CMD ["/legit-node"]