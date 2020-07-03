FROM golang:1.14.4-alpine3.12 AS builder

WORKDIR /go/src/puke

COPY . .

RUN go install

FROM alpine:3.12.0

COPY --from=builder /go/bin/puke /usr/bin/puke
RUN chmod +x /usr/bin/puke

ENTRYPOINT ["/usr/bin/puke"]
