FROM golang:1.16-alpine3.12 as builder

RUN go version

RUN apk update && \
  apk add --no-cache git ca-certificates tzdata build-base make && \
  update-ca-certificates

COPY . /src/

WORKDIR /src

RUN cd /src && \
  go mod download

RUN cd /src && \
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -ldflags "-s -w" -o /opt/app/app && \
  chmod +x /opt/app/app


FROM alpine:3.12

RUN apk add --no-cache curl ca-certificates

WORKDIR /opt/app

COPY --from=builder /opt/app/app /opt/app/app
COPY --from=builder /src/docker-entrypoint.sh /
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

ENV ZONEINFO=/zoneinfo.zip

RUN mkdir /opt/app/uploads && \
  chmod +x /opt/app/app && \
  chmod +x /docker-entrypoint.sh

EXPOSE 8083

ENTRYPOINT [ "/docker-entrypoint.sh", "/opt/app/app" ]
