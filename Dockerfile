FROM golang:1.19-alpine3.17 as go-build
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
RUN apk add git
# golang配置
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
# 程序入口
ARG CMD_PATH=./cmd/midjourney-apiserver
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app $CMD_PATH

FROM alpine:3.17 as production
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories
# 设置时区为上海的时区
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
WORKDIR /app
COPY --from=go-build /app/app .
# 复制资源文件
COPY /conf/conf.yml /conf/conf.yml
EXPOSE 8080
CMD ["./app"]
