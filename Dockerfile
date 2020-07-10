FROM golang:1.12

WORKDIR /go/src/github.com/tokend/hgate
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-X github.com/tokend/hgate/internal/config.HGateRelease=${VERSION}" \
    -o /usr/local/bin/hgate github.com/tokend/hgate

###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/hgate /usr/local/bin/hgate
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["hgate"]
