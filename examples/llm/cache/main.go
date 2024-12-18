package main

import (
	"context"
	"fmt"

	// openaiembedder "github.com/henomis/lingoose/embedder/openai"
	ollamaembedder "github.com/henomis/lingoose/embedder/ollama"
	"github.com/henomis/lingoose/index"
	"github.com/henomis/lingoose/index/vectordb/jsondb"
	"github.com/henomis/lingoose/llm/cache"

	// "github.com/henomis/lingoose/llm/openai"
	"github.com/henomis/lingoose/llm/ollama"
	"github.com/henomis/lingoose/thread"
)

func main() {
	/*
		llm := openai.New().WithCache(
			cache.New(
				index.New(
					jsondb.New().WithPersist("index.json"),
					openaiembedder.New(openaiembedder.AdaEmbeddingV2),
				),
			).WithTopK(3),
		)*/

	llm := ollama.New().WithCache(
		cache.New(
			index.New(
				jsondb.New().WithPersist("index.json"),
				ollamaembedder.New(),
			)).WithTopK(3),
	).WithStream(func(s string) {
		fmt.Print(s)
	})

	questions := []string{
		"what's github",
		//		"can you explain what GitHub is",
		//		"can you tell me more about GitHub",
		"what is the purpose of GitHub",
		"what is the Github?",
		"What is Github",
	}

	for _, question := range questions {
		t := thread.New().AddMessage(
			thread.NewUserMessage().AddContent(
				thread.NewTextContent(question),
			),
		)

		err := llm.Generate(context.Background(), t)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(t)
	}
}
