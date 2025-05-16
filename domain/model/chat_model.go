package model

// ChatMessage represents a single message in a chat conversation
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ChatSession represents a chat conversation
type ChatSession struct {
	Messages []ChatMessage
}

// NewChatSession creates a new chat session
func NewChatSession() *ChatSession {
	return &ChatSession{
		Messages: make([]ChatMessage, 0),
	}
}

// AddMessage adds a new message to the chat
func (c *ChatSession) AddMessage(role, content string) {
	c.Messages = append(c.Messages, ChatMessage{
		Role:    role,
		Content: content,
	})
}
