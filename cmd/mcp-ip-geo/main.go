package main

import (
	"flag"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/server"
)

func main() {
	addr := flag.String("address", "", "The host and port to run the sse server")
	flag.Parse()

	if err := server.Run(*addr); err != nil {
		panic(err)
	}
}
