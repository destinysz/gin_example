FROM golang:alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.io,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 编译成二进制可执行文件app
RUN go build -o app .

# 超小镜像
FROM scratch

# 从builder镜像中把app 拷贝到当前目录
COPY --from=builder /build/app /

COPY ./config /config

# 需要运行的命令
ENTRYPOINT ["./app"]