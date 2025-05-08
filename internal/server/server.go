package server

import (
	"errors"
	"log"

	"github.com/ThinkInAIXYZ/go-mcp/server"
	"github.com/ThinkInAIXYZ/go-mcp/transport"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/domain"
	"github.com/chenmingyong0423/mcp-ip-geo/internal/tools"
)

func Run(address string, transportProtocol domain.TransportProtocol) error {
	var err error

	var ts transport.ServerTransport

	switch transportProtocol {
	case domain.Stdio:
		ts = transport.NewStdioServerTransport()
	case domain.Sse:
		ts, err = transport.NewSSEServerTransport(address)
		if err != nil {
			return err
		}
	case domain.StreamableHttp:
		log.Printf("run streamable http server")
		ts = transport.NewStreamableHTTPServerTransport(address)
	default:
		return errors.New("invalid transport protocol")
	}

	s, err := server.NewServer(ts)
	if err != nil {
		return err
	}

	toolFuncList := tools.GetToolFuncList()
	for _, tool := range toolFuncList {
		s.RegisterTool(tool())
	}

	return s.Run()
}
