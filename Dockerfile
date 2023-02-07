FROM golang:latest as go
WORKDIR /src
COPY ./ /src
RUN go env -w GO111MODULE=on; \
    go env -w GOPROXY=https://goproxy.cn,direct
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./web

###
FROM alpine:latest
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk update --no-cache && \
    apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=go /src/main /app

CMD [ "/app/main" ]