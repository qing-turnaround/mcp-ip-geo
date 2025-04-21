# 
<h1 align="center">
  mcp-ip-geo
</h1>

---

`mcp-ip-geo` is an MCP Server that provides IP geolocation lookup (country, region, city, etc.) via ip-api.com.

# Build from source

- local

```bash
 go build -o mcp-ip-geo .\cmd\mcp-ip-geo
 
 go build -o mcp-ip-geo.exe .\cmd\mcp-ip-geo # windows
```

- remote

```bash
 go install github.com/chenmingyong0423/mcp-ip-geo/cmd/mcp-ip-geo@latest
```