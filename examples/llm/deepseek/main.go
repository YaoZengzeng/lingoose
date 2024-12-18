package main

import (
	"context"
	"fmt"

	"github.com/henomis/lingoose/llm/deepseek"
	"github.com/henomis/lingoose/thread"
)

func main() {
	// The DeepSeek API key is expected to be set in the DEEPSEEK_API_KEY environment variable
	dsllm := deepseek.New().WithModel("deepseek-chat")

	t := thread.New().AddMessage(
		thread.NewUserMessage().AddContent(
			thread.NewTextContent("What's the NATO purpose?"),
		),
	)

	err := dsllm.Generate(context.Background(), t)
	if err != nil {
		panic(err)
	}

	fmt.Println(t)
}
