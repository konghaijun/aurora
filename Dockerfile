# 设置基础镜像为 alpine
FROM alpine:latest

# 在容器内创建一个工作目录
WORKDIR /app

# 复制二进制文件到容器内的工作目录
COPY douyin /app/douyin

# 设置可执行权限
RUN chmod +x /app/douyin

# 暴露端口（如果需要）
EXPOSE 8080

# 定义容器启动时的入口点
CMD ["/app/douyin"]