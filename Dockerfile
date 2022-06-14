# syntax=docker/dockerfile:1
FROM golang:1.18.3-alpine
WORKDIR $GOPATH/src/payment-service/
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/payment-service ./cmd
COPY /config /go/bin/config

FROM alpine
COPY --from=0 /go/bin/config /go/bin/config
COPY --from=0 /go/bin/payment-service /go/bin/payment-service
EXPOSE 8080
ENTRYPOINT ["/go/bin/payment-service"]