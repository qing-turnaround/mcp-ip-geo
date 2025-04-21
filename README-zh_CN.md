<h1 align="center">
  mcp-ip-geo
</h1>

<div align="center">
  <a href="./README.md">English</a>
</div>

---

`mcp-ip-geo` 是一个 MCP 服务器，提供通过 ip-api.com 实现的 IP 地理位置查询功能（国家、地区、城市等信息）。

# 从源码构建
## 本地模式
### 使用 Go 命令
```bash
go build -o mcp-ip-geo .\cmd\mcp-ip-geo
 
go build -o mcp-ip-geo.exe .\cmd\mcp-ip-geo # Windows 系统
```

### 使用 Docker
- 1. 构建 Docker 镜像
  ```bash
  docker build -t mcp-ip-geo-server .
  ```
- 2. 运行 Docker 容器
  ```bash
  docker run -d --name mcp-ip-geo-server -p 8000:8000 mcp-ip-geo-server
  ```

### 远程模式
使用 Go 安装服务端：
```bash
go install github.com/chenmingyong0423/mcp-ip-geo/cmd/mcp-ip-geo@latest
```