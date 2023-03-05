FROM golang:1.19-alpine as builder
ENV GOBIN=/go/src/github.com/toyohashi6140/doscord-bot/bin
ENV PATH ${GOBIN}:$PATH

WORKDIR /go/src/github.com/toyohashi6140/doscord-bot

COPY go.mod .
COPY go.sum .

RUN apk add --no-cache git alpine-sdk
RUN set -x && go mod download

COPY . .
RUN go build -o /main ./main.go

FROM alpine:latest
EXPOSE 8080
RUN apk update && apk upgrade && \
    apk add ca-certificates && rm -rf /var/cache/apk/* && \
    apk add tzdata && mkdir -m 666 -p /tmp
COPY --from=builder /main ./app/main
USER nobody
CMD [ "./app/main" ]

