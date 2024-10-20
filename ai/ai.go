package ai

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
	log "github.com/sirupsen/logrus"
)

type AI struct {
	client *openai.Client
}

func NewAI() *AI {
	return &AI{
		client: getClient(),
	}
}

func getClient() *openai.Client {
	key := ""

	cfg := openai.DefaultConfig(key)
	cfg.BaseURL = "https://api-us.luee.net/v1"

	client := openai.NewClientWithConfig(cfg)
	return client
}

func (ai *AI) NewSession(ctx context.Context, prompt string) *Session {
	s := &Session{
		client: ai.client,
		ctx:    ctx,
		Msg: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt,
			},
		},
		Results: make([]openai.ChatCompletionMessage, 0),
	}

	return s
}

func (ai *AI) Hello() string {
	resp, err := ai.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)

	if err != nil {
		log.WithField("error", err).Error("Create ChatCompletion error")
		return ""
	}

	log.WithField("> ", resp.Choices[0].Message.Content).Info("ai.Hello")
	return resp.Choices[0].Message.Content
}

func (ai *AI) StreamMsg() {
	ctx := context.Background()
	type Result struct {
		Steps []struct {
			Explanation string `json:"explanation"`
			Output      string `json:"output"`
		} `json:"steps"`
		FinalAnswer string `json:"final_answer"`
	}
	var result Result
	schema, err := jsonschema.GenerateSchemaForType(result)
	if err != nil {
		log.Fatalf("GenerateSchemaForType error: %v", err)
	}
	resp, err := ai.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are a helpful math tutor. Guide the user through the solution step by step.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "how can I solve 8x + 7 = -23",
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:   "math_reasoning",
				Schema: schema,
				Strict: true,
			},
		},
	})
	if err != nil {
		log.WithError(err).Error("CreateChatCompletion error")
		return
	}
	err = schema.Unmarshal(resp.Choices[0].Message.Content, &result)
	if err != nil {
		log.WithError(err).Error("Unmarshal schema error")
	}
	// loop print result steps.
	for i, step := range result.Steps {
		log.WithField("step", i).WithField("todo", step).Info("ai.StreamMsg")
	}
	log.WithField("final answer", result.FinalAnswer)
}
