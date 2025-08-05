package chatGPT

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/project-flogo/core/activity"
)

var activityMetadata = NewMetadata()

type ChatRequest struct {
	Model    string      `json:"model"`
	Messages []ChatEntry `json:"messages"`
}

type ChatEntry struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Choices []struct {
		Message ChatEntry `json:"message"`
	} `json:"choices"`
}

// Activity is a ChatGPT API activity
type Activity struct{}

func New(ctx context.Context, settings map[string]interface{}) (activity.Activity, error) {
	return &Activity{}, nil
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMetadata
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	apiKey := ctx.GetInput("apiKey").(string)
	model := ctx.GetInput("model").(string)
	prompt := ctx.GetInput("prompt").(string)

	reqBody := ChatRequest{
		Model: model,
		Messages: []ChatEntry{
			{Role: "user", Content: prompt},
		},
	}

	bodyBytes, _ := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(bodyBytes))
	if err != nil {
		return false, err
	}

	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("OpenAI API error: %s", resp.Status)
	}

	respBody, _ := io.ReadAll(resp.Body)

	var chatResp ChatResponse
	err = json.Unmarshal(respBody, &chatResp)
	if err != nil {
		return false, err
	}

	if len(chatResp.Choices) == 0 {
		return false, fmt.Errorf("no response from ChatGPT")
	}

	ctx.SetOutput("response", chatResp.Choices[0].Message.Content)
	return true, nil
}
