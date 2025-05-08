package tools

import (
	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/ThinkInAIXYZ/go-mcp/server"
)

type ToolFunc func() (tool *protocol.Tool, toolHandler server.ToolHandlerFunc)

func GetToolFuncList() []ToolFunc {
	return []ToolFunc{
		// SingleIpParser,
		TestA0Tool,
		TestA1Tool,
		TestA2Tool,
	}
}

var (
	singleIpParserTool *protocol.Tool
	testA0Tool         *protocol.Tool
	testA1Tool         *protocol.Tool
	testA2Tool         *protocol.Tool
)

func init() {
	var err error
	// singleIpParserTool, err = protocol.NewTool("ip-details", "a tool that provides IP geolocation information", ipRequest{})
	// if err != nil {
	// 	panic(err)
	// }

	testA0Tool, err = protocol.NewTool("test-a0", "可以创建一个资源目录", testA0Request{})
	if err != nil {
		panic(err)
	}
	testA1Tool, err = protocol.NewTool("test-a1", "可以创建一个资源文件", testA1Request{})
	if err != nil {
		panic(err)
	}
	testA2Tool, err = protocol.NewTool("test-a2", "可以修改资源文件内容", testA2Request{})
	if err != nil {
		panic(err)
	}
}
