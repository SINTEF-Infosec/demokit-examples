FROM golang:1.16 as build
WORKDIR /go/src/light-controller/
COPY . .
RUN go build -o /build/light-controller /go/src/light-controller

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/light-controller /

CMD ["/light-controller"]