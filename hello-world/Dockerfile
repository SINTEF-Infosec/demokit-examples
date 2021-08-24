FROM golang:1.16 as build
WORKDIR /go/src/hello-world/
COPY . .
RUN go build -o /build/hello-world /go/src/hello-world

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/hello-world /

CMD ["/hello-world"]