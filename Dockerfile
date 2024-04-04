FROM golang:1.22-alpine AS builder

WORKDIR /build

COPY . .

RUN apk update --no-cache && \
    apk add --no-cache make nodejs npm

RUN npm install -g pnpm

RUN make

FROM alpine:latest

COPY --from=builder /build/bin/reblog /app/reblog

WORKDIR /root

CMD ["/app/reblog"]
