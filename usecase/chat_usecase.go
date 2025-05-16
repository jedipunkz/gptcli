package usecase

import (
	"gptcli/domain/model"
	"gptcli/domain/repository"
)

type ChatUseCase struct {
	repo repository.ChatRepository
}

func NewChatUseCase(repo repository.ChatRepository) *ChatUseCase {
	return &ChatUseCase{
		repo: repo,
	}
}

func (u *ChatUseCase) StartChat(modelName string) *model.ChatSession {
	return model.NewChatSession()
}

func (u *ChatUseCase) SendMessage(chat *model.ChatSession, content string, modelName string) (*model.ChatMessage, error) {
	chat.AddMessage("user", content)
	response, err := u.repo.CreateChatCompletion(chat, modelName)
	if err != nil {
		return nil, err
	}
	chat.AddMessage(response.Role, response.Content)
	return response, nil
}
