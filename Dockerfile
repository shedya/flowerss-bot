FROM golang:1.22.2-alpine as builder
#ENV CGO_ENABLED=0
COPY . /flowerss
# 添加gcompat包以提供glibc兼容性
RUN apk add --no-cache gcompat
# 安装必要的构建工具和库
RUN apk add --no-cache git make gcc libc-dev&& \
    cd /flowerss && make build

# Image starts here
FROM alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /flowerss/flowerss-bot /bin/
VOLUME /root/.flowerss
WORKDIR /root/.flowerss
ENTRYPOINT ["/bin/flowerss-bot"]

