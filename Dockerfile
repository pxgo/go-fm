# syntax=docker/dockerfile:1.7
# ------------------------------
# 构建阶段：编译 Go 二进制
# ------------------------------
FROM golang:1.20-alpine AS builder

WORKDIR /src

# 仅拷贝依赖声明，加速缓存
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# 拷贝剩余源码
COPY . .

# 版本号（可通过 --build-arg VERSION=xxx 传入，不传则从 conf/version.txt 读取）
ARG VERSION_FILE=conf/version.txt
ARG VERSION
RUN if [ -z "$VERSION" ] && [ -f "$VERSION_FILE" ]; then VERSION=$(cat $VERSION_FILE); fi && \
    echo "> Build version: $VERSION" && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o /out/gofm .

# ------------------------------
# 运行阶段：极简镜像 (alpine)
# 如需 scratch/distroless 可再精简
# ------------------------------
FROM alpine:3.20 AS runtime

# 创建非 root 用户
RUN adduser -D -H -u 10001 appuser
WORKDIR /app

# 拷贝二进制
COPY --from=builder /out/gofm /app/gofm

# 运行时可挂载音乐目录
VOLUME ["/music"]

# 默认监听端口（与程序默认一致，可通过 -p 改）
EXPOSE 8090

USER appuser

# 入口：可再追加参数
ENTRYPOINT ["/app/gofm"]
# 默认参数：播放 /music 目录
CMD ["-d", "/music"]
