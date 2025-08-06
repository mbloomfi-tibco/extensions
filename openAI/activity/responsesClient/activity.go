package responsesClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/project-flogo/core/activity"
)

// activityMd is the metadata for the activity.
var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// Metadata returns the activity's metadata.
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

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
type Activity struct {
	outputFormat string
}

// New creates a new instance of the Activity.
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := s.FromMap(ctx.Settings())
	if err != nil {
		return nil, err
	}

	act := &Activity{
		outputFormat: s.OutputFormat,
	}
	return act, nil
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {
	apiKey := ctx.GetInput(iAPIKey).(string)
	model := ctx.GetInput(iModel).(string)
	prompt := ctx.GetInput(iPrompt).(string)

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

	ctx.SetOutput(oResponse, chatResp.Choices[0].Message.Content)
	return true, nil
}
