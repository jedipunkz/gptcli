package model

// GenerationRequest represents a text generation request
type GenerationRequest struct {
	Prompt      string  `json:"prompt"`
	Temperature float32 `json:"temperature"`
	MaxTokens   int     `json:"max_tokens"`
}

// NewGenerationRequest creates a new generation request
func NewGenerationRequest(prompt string, temperature float32, maxTokens int) *GenerationRequest {
	return &GenerationRequest{
		Prompt:      prompt,
		Temperature: temperature,
		MaxTokens:   maxTokens,
	}
}
