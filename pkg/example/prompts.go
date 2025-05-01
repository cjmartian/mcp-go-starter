package example

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func NewHelloPrompt() (mcp.Prompt, server.PromptHandlerFunc) {
	prompt := mcp.NewPrompt("hello_prompt",
		mcp.WithPromptDescription("A prompt that greets the user by name."),
		mcp.WithArgument("name",
			mcp.ArgumentDescription("Name of the person to greet"),
		),
	)
	handler := func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		name := req.Params.Arguments["name"]
		if name == "" {
			name = "friend"
		}
		return mcp.NewGetPromptResult(
			"A friendly greeting",
			[]mcp.PromptMessage{
				mcp.NewPromptMessage(
					mcp.RoleAssistant,
					mcp.NewTextContent(fmt.Sprintf("Hello, %s! How can I help you today?", name)),
				),
			},
		), nil
	}
	return prompt, handler
}

func NewWikiSummaryPrompt() (mcp.Prompt, server.PromptHandlerFunc) {
	prompt := mcp.NewPrompt("wiki_summary_prompt",
		mcp.WithPromptDescription("Fetch a summary from Wikipedia"),
		mcp.WithArgument("page",
			mcp.ArgumentDescription("The page to fetch a summary for"),
		),
	)
	handler := func(ctx context.Context, req mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		page := req.Params.Arguments["page"]
		if page == "" {
			return nil, fmt.Errorf("page argument is required")
		}
		return mcp.NewGetPromptResult(
			"Fetching summary",
			[]mcp.PromptMessage{
				mcp.NewPromptMessage(
					mcp.RoleAssistant,
					mcp.NewTextContent(fmt.Sprintf("Fetching summary for: %s", page)),
				),
			},
		), nil
	}
	return prompt, handler
}

func RegisterPrompts(s *server.MCPServer) {
	prompt, handler := NewHelloPrompt()
	s.AddPrompt(prompt, handler)

	// wikiPrompt, wikiHandler := NewWikiSummaryPrompt()
	// s.AddPrompt(wikiPrompt, wikiHandler)
}
