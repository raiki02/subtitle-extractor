package video_summary

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/components/tool/utils"
)

type SummarizeInput struct {
	VideoPath    string `json:"video_path" jsonschema:"required" jsonschema_description:"Local video file path to summarize."`
	Prompt       string `json:"prompt,omitempty" jsonschema_description:"Optional Marlin caption prompt override. Defaults to the configured prompt."`
	MaxNewTokens int    `json:"max_new_tokens,omitempty" jsonschema_description:"Maximum number of generated tokens. Defaults to the configured value."`
}

type FindInput struct {
	VideoPath    string `json:"video_path" jsonschema:"required" jsonschema_description:"Local video file path to inspect."`
	Event        string `json:"event" jsonschema:"required" jsonschema_description:"Natural-language event to locate in the video."`
	MaxNewTokens int    `json:"max_new_tokens,omitempty" jsonschema_description:"Maximum number of generated tokens. Defaults to the configured value."`
}

func NewSummarizeTool(client *Client) (tool.InvokableTool, error) {
	return utils.InferTool(
		"summarize_video",
		"Summarize a local video by calling the NemoStation Marlin video understanding service.",
		func(ctx context.Context, input SummarizeInput) (CaptionResponse, error) {
			return client.Caption(ctx, input.VideoPath, SummarizeOptions{
				Prompt:       input.Prompt,
				MaxNewTokens: input.MaxNewTokens,
			})
		},
	)
}

func NewFindTool(client *Client) (tool.InvokableTool, error) {
	return utils.InferTool(
		"find_video_event",
		"Find the time span of an event in a local video by calling the NemoStation Marlin video understanding service.",
		func(ctx context.Context, input FindInput) (FindResponse, error) {
			return client.Find(ctx, input.VideoPath, input.Event, SummarizeOptions{
				MaxNewTokens: input.MaxNewTokens,
			})
		},
	)
}
