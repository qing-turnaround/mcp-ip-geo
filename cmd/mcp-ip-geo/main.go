package main

import (
	"flag"

	"github.com/chenmingyong0423/mcp-ip-geo/internal/domain"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/server"
)

var serverTransport = domain.StreamableHttp

func main() {
	addr := flag.String("address", ":8080", "The host and port to run the sse server")
	flag.Parse()

	if addr == nil {
		serverTransport = domain.Stdio
	}

	if err := server.Run(*addr, serverTransport); err != nil {
		panic(err)
	}
}
