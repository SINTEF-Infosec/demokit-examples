FROM golang:1.16 as build
WORKDIR /go/src/simple-greeter/
COPY . .
RUN go build -o /build/simple-greeter /go/src/simple-greeter

# The production image, based on google distroless
# See https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10

COPY --from=build /build/simple-greeter /

CMD ["/simple-greeter"]