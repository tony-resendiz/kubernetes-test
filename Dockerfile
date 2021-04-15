FROM golang:alpine AS builder
WORKDIR /
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates
COPY src/main/webserver.go .
RUN go mod init main
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/webserver

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/webserver /go/bin/webserver
ENTRYPOINT [ "/go/bin/webserver" ] 