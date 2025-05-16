package usecase

import (
	"gptcli/domain/repository"
)

type GenerationUseCase struct {
	repo repository.GenerationRepository
}

func NewGenerationUseCase(repo repository.GenerationRepository) *GenerationUseCase {
	return &GenerationUseCase{
		repo: repo,
	}
}

func (u *GenerationUseCase) GenerateText(prompt string, modelName string, temperature float32, maxTokens int, stream bool) (string, error) {
	return u.repo.CreateCompletion(prompt, modelName, temperature, maxTokens, stream)
}
