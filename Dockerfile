FROM golang:1.22.2-alpine3.19 AS builder

ARG CNMIRROR=false

WORKDIR /app

COPY . .

RUN if [ "$CNMIRROR" = "true" ]; then \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories; \
fi

RUN apk update --no-cache && \
    apk add --no-cache make nodejs npm gcc musl-dev

RUN if [ "$CNMIRROR" = "true" ]; then \
    npm config set registry https://registry.npmmirror.com/; \
fi


RUN if [ "$CNMIRROR" = "true" ]; then go env -w GOPROXY=https://goproxy.cn,direct; fi

# Fuck CGO
RUN go env -w CGO_ENABLED=0

RUN make install-dev

RUN make

FROM alpine:latest

COPY --from=builder /app/bin/reblog /app/reblog

WORKDIR /root

CMD ["/app/reblog"]

EXPOSE 3000
