<<<<<<< HEAD
<<<<<<< HEAD
FROM golang:1.20.12-alpine as builder
=======
FROM golang:1.20.10-alpine as builder
>>>>>>> parent of ba607ec (back 1.18)
=======
FROM golang:1.20.10-alpine as builder
>>>>>>> parent of ba607ec (back 1.18)
#ENV CGO_ENABLED=0
COPY . /flowerss
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

