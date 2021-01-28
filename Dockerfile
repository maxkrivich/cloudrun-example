# Base builder stage
FROM golang:1.15-buster as build
WORKDIR /go/src/app
ADD . /go/src/app
RUN go get -d -v ./...
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-linkmode external -extldflags "-static"' -o /go/bin/app .

# Final stage with the binnary file
FROM gcr.io/distroless/base
COPY --from=build /go/bin/app /
CMD ["/app"]