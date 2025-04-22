<h1 align="center">
  mcp-ip-geo
</h1>

<div align="center">
  <a href="./README-zh_CN.md">简体中文</a>
</div>

---

`mcp-ip-geo` is an MCP server that provides IP geolocation lookup (country, region, city, etc.) using the ip-api.com service.

# Building from Source

## Build Locally

### Using Go Command

```bash
# On Unix-based systems (Linux/macOS)
go build -o mcp-ip-geo ./cmd/mcp-ip-geo

# On Windows
go build -o mcp-ip-geo.exe .\cmd\mcp-ip-geo
```

### Using Docker

1. Build the Docker image:

    ```bash
    docker build -t mcp-ip-geo-server .
    ```

2. Run the Docker container:

    ```bash
    docker run -d --name mcp-ip-geo-server -p 8000:8000 mcp-ip-geo-server
    ```

## Install Prebuilt Binary

Install the server using Go:

```bash
go install github.com/chenmingyong0423/mcp-ip-geo/cmd/mcp-ip-geo@latest
```

# Integrating with MCP

You can integrate `mcp-ip-geo` with the MCP system in one of the following ways:

## 🖥 Executable Integration (run a local binary)

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "command": "/path/to/mcp-ip-geo"
    }
  }
}
```

## 🌐 HTTP Integration (connect to a running instance via HTTP)

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "url": "http://host:port/sse"
    }
  }
}
```
