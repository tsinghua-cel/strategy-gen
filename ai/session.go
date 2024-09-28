package ai

import (
	"context"
	"github.com/sashabaranov/go-openai"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Session struct {
	ctx     context.Context
	client  *openai.Client
	Msg     []openai.ChatCompletionMessage
	Results []openai.ChatCompletionMessage
	mux     sync.Mutex
}

func (s *Session) Ask(msg string) (string, error) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.Msg = append(s.Msg, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: msg,
	})
	resp, err := s.client.CreateChatCompletion(
		s.ctx,
		openai.ChatCompletionRequest{
			Model:    openai.GPT4oMini,
			Messages: s.Msg,
		},
	)

	if err != nil {
		log.WithField("error", err).Error("Create ChatCompletion error")
		return "", err
	}
	s.Msg = append(s.Msg, resp.Choices[0].Message)
	s.Results = append(s.Results, resp.Choices[0].Message)
	return s.Results[len(s.Results)-1].Content, nil
}
