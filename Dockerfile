FROM golang:1.20.14-alpine as builder
#ENV CGO_ENABLED=0
COPY . /flowerss
# 安装必要的构建工具和库
RUN apk add git make gcc libc-dev && \
    cd /flowerss && make build

# Image starts here
FROM alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /flowerss/flowerss-bot /bin/
VOLUME /root/.flowerss
WORKDIR /root/.flowerss
ENTRYPOINT ["/bin/flowerss-bot"]

