package tools

import (
	"context"
	"fmt"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
)

type testA0Request struct {
	FolderName string `json:"folderName" description:"创建目录的名字" required:"true"`
}

type testA1Request struct {
	FolderName string `json:"folderName" description:"创建文件的目录" required:"true"`
	FileName   string `json:"fileName" description:"创建文件的名字" required:"true"`
}

type testA2Request struct {
	Content string `json:"content" description:"修改文件的内容" required:"true"`
}

type commonRequest struct {
	Result string `json:"result" description:"结果" required:"true"`
}

func TestA0Tool() (*protocol.Tool, server.ToolHandlerFunc) {
	return testA0Tool, func(ctx context.Context, toolRequest *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
		var req testA0Request
		if err := protocol.VerifyAndUnmarshal(toolRequest.RawArguments, &req); err != nil {
			return nil, err
		}
		fmt.Println(req)
		return protocol.NewCallToolResult([]protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: "成功",
			},
		}, false), nil
	}
}

func TestA1Tool() (*protocol.Tool, server.ToolHandlerFunc) {
	return testA1Tool, func(ctx context.Context, toolRequest *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
		var req testA1Request
		if err := protocol.VerifyAndUnmarshal(toolRequest.RawArguments, &req); err != nil {
			return nil, err
		}
		fmt.Println(req)
		return protocol.NewCallToolResult([]protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: "成功",
			},
		}, false), nil
	}
}

func TestA2Tool() (*protocol.Tool, server.ToolHandlerFunc) {
	return testA2Tool, func(ctx context.Context, toolRequest *protocol.CallToolRequest) (*protocol.CallToolResult, error) {
		var req testA2Request
		if err := protocol.VerifyAndUnmarshal(toolRequest.RawArguments, &req); err != nil {
			return nil, err
		}
		fmt.Println(req)
		return protocol.NewCallToolResult([]protocol.Content{
			protocol.TextContent{
				Type: "text",
				Text: "成功",
			},
		}, false), nil
	}
}
