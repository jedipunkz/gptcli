package repository

import "gptcli/domain/model"

// ChatRepository defines the interface for chat operations
type ChatRepository interface {
	// CreateChatCompletion sends messages to the AI and returns the response
	CreateChatCompletion(chat *model.ChatSession, modelName string) (*model.ChatMessage, error)
}
