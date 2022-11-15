# 编译镜像
FROM golang:1.19 as build

# 只要go.mod,go.sum,vendor内容不变,依赖会使用缓存
# 否则会全部重新下载
WORKDIR /go/cache
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

# 编译代码
WORKDIR /go/release
ADD . .

# -s: 省略符号表和调试信息
# -w: 省略DWARF符号表
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o srv-bin ./cmd/server/

# 运行镜像
 FROM alpine:3.12

WORKDIR /app

COPY --from=build /go/release/srv-bin /app/
COPY config/i18n /app/conf/i18n
COPY api/mozi/device/v1/*.proto /app/conf/proto/

RUN mkdir -p /app/conf

ENV CONF_FILE="/app/conf/app.yaml"

# http port
#EXPOSE 8102
# grpc port
#EXPOSE 9102

CMD ["./srv-bin", "-conf", "/app/conf/app.yaml"]
