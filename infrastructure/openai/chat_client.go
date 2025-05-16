package openai

import (
	"context"
	"gptcli/domain/model"
	"gptcli/domain/repository"

	openai "github.com/sashabaranov/go-openai"
)

type ChatClient struct {
	client *openai.Client
}

// NewChatClient creates a new OpenAI chat client
func NewChatClient(apiKey string) repository.ChatRepository {
	return &ChatClient{
		client: openai.NewClient(apiKey),
	}
}

// CreateChatCompletion implements the ChatRepository interface
func (c *ChatClient) CreateChatCompletion(chat *model.ChatSession, modelName string) (*model.ChatMessage, error) {
	messages := make([]openai.ChatCompletionMessage, len(chat.Messages))
	for i, msg := range chat.Messages {
		messages[i] = openai.ChatCompletionMessage{
			Role:    msg.Role,
			Content: msg.Content,
		}
	}

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    modelName,
			Messages: messages,
		},
	)
	if err != nil {
		return nil, err
	}

	return &model.ChatMessage{
		Role:    "assistant",
		Content: resp.Choices[0].Message.Content,
	}, nil
}
