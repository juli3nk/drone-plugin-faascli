FROM golang:1.14-alpine3.11 AS builder

RUN apk --update add \
		ca-certificates \
		gcc \
		git \
		musl-dev

COPY go.mod go.sum /go/src/github.com/juli3nk/drone-plugin-faascli/
WORKDIR /go/src/github.com/juli3nk/drone-plugin-faascli

ENV GO111MODULE on
RUN go mod download

COPY . .

RUN go build -ldflags "-linkmode external -extldflags -static -s -w" -o /tmp/drone-faascli


FROM openfaas/faas-cli as faascli


FROM docker:19.03-dind

RUN apk --update --no-cache add \
		git

COPY --from=builder /tmp/drone-faascli /usr/local/bin/
COPY --from=faascli /usr/bin/faas-cli /usr/local/bin/

ENTRYPOINT [ "/usr/local/bin/drone-faascli" ]
