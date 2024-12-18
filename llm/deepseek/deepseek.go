package deepseek

import (
	"os"

	"github.com/henomis/lingoose/llm/openai"
	goopenai "github.com/sashabaranov/go-openai"
)

const (
	deepseekAPIEndpoint = "https://api.deepseek.com/"
)

type DeepSeek struct {
	*openai.OpenAI
}

func New() *DeepSeek {
	customConfig := goopenai.DefaultConfig(os.Getenv("DEEPSEEK_API_KEY"))
	customConfig.BaseURL = deepseekAPIEndpoint
	customClient := goopenai.NewClientWithConfig(customConfig)

	openaillm := openai.New().WithClient(customClient)
	openaillm.Name = "deepseek"

	return &DeepSeek{
		OpenAI: openaillm,
	}
}
