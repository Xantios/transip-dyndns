FROM golang:1.21.1-alpine3.18 as compiler

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY * /app
RUN go build -o /dyndns

FROM scratch

COPY --from=compiler /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=compiler /dyndns /dyndns
COPY ./key.private /key.private

ENTRYPOINT ["/dyndns"]