FROM golang:1.16 as build
WORKDIR /go/src/token-ring/
COPY . .
RUN go build -o /build/token-ring /go/src/token-ring

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/token-ring /

CMD ["/token-ring"]