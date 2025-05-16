package openai

import (
	"context"
	"fmt"
	"gptcli/domain/repository"
	"io"

	openai "github.com/sashabaranov/go-openai"
)

type GenerationClient struct {
	client *openai.Client
}

// NewGenerationClient creates a new OpenAI generation client
func NewGenerationClient(apiKey string) repository.GenerationRepository {
	return &GenerationClient{
		client: openai.NewClient(apiKey),
	}
}

// CreateCompletion implements the GenerationRepository interface
func (c *GenerationClient) CreateCompletion(prompt string, modelName string, temperature float32, maxTokens int, useStream bool) (string, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	if !useStream {
		resp, err := c.client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:       modelName,
				Messages:    messages,
				Temperature: temperature,
				MaxTokens:   maxTokens,
			},
		)
		if err != nil {
			return "", err
		}
		return resp.Choices[0].Message.Content, nil
	}

	// ストリーミングモード
	stream, err := c.client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       modelName,
			Messages:    messages,
			Temperature: temperature,
			MaxTokens:   maxTokens,
		},
	)
	if err != nil {
		return "", err
	}
	defer stream.Close()

	var fullResponse string
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}

		content := response.Choices[0].Delta.Content
		fmt.Print(content)
		fullResponse += content
	}
	fmt.Println()
	return fullResponse, nil
}
