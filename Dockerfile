FROM golang:1.12.5 AS builder

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

ADD . /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO11MODULE=on go build -v -a -o /main .

FROM alpine:3.7
COPY --from=builder /main /kubernetes-event-exporter
ENTRYPOINT ["/kubernetes-event-exporter"]