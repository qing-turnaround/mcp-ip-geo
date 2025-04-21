<h1 align="center">
  mcp-ip-geo
</h1>

<div align="center">
  <a href="./README-zh_CN.md">中文简体</a>
</div>

---

`mcp-ip-geo` is an MCP Server that provides IP geolocation lookup (country, region, city, etc.) via ip-api.com.

# Build from source
## Local mode
### Go Command
```bash
go build -o mcp-ip-geo .\cmd\mcp-ip-geo
 
go build -o mcp-ip-geo.exe .\cmd\mcp-ip-geo # windows
```
### Docker
- 1. Build the Docker image
    ```bash
    docker build -t mcp-ip-geo-server .
    ```
- 2. Run the Docker container
    ```bash
    docker run -d --name mcp-ip-geo-server -p 8000:8000 mcp-ip-geo-server
    ```

### Remote mode
Install the server using Go:
```bash
 go install github.com/chenmingyong0423/mcp-ip-geo/cmd/mcp-ip-geo@latest
```
# MCP Server Config
- Local File

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "command": "/path/to/mcp-ip-geo"
    }
  }
}
```

- `Remote mode`

```json
{
  "mcpServers": {
    "mcp-ip-geo": {
      "url": "http://host:port/sse"
    }
  }
}
```