FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 复制编译好的二进制文件
COPY cmd/zabbix-api-server/zabbix-api-server .

# 设置可执行权限
RUN chmod +x ./zabbix-api-server

# 暴露端口
EXPOSE 8080

# 默认命令
CMD ["./zabbix-api-server"]
