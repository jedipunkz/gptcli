package repository

// GenerationRepository defines the interface for text generation operations
type GenerationRepository interface {
	// CreateCompletion generates text based on a prompt
	CreateCompletion(prompt string, modelName string, temperature float32, maxTokens int, stream bool) (string, error)
}
