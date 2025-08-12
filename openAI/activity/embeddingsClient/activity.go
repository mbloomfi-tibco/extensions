package embeddingsClient

/*
* Copyright © 2023 - 2024. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"context"
	"log"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"

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

// Activity is a ChatGPT API activity
type Activity struct {
	apiKey       string
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
		apiKey:       s.ApiKey,
		outputFormat: s.OutputFormat,
	}

	log.Printf("Activity initialized with API Key: %s and Output Format: %s", act.apiKey, act.outputFormat)

	return act, nil
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	//model := ctx.GetInput(iModel).(string)
	embeddingText := ctx.GetInput(iPrompt).(string)
	//tool := ctx.GetInput(iTool).(string)

	if a.apiKey == "" {
		log.Fatal("Missing openAPI key")
	}

	oaiClient := openai.NewClient(
		option.WithAPIKey(a.apiKey),
	)

	resp, err := oaiClient.Embeddings.New(
		context.Background(),
		openai.EmbeddingNewParams{
			Model: openai.EmbeddingModelTextEmbedding3Small, // or openai.TextEmbedding3Large
			Input: openai.EmbeddingNewParamsInputUnion{
				OfString: openai.String(embeddingText),
			},
			EncodingFormat: openai.EmbeddingNewParamsEncodingFormatFloat,
		},
	)

	if err != nil {
		log.Fatalf("Embedding creation error: %v", err)
	}

	// Print vector length and first few values
	vector := resp.Data[0].Embedding
	//fmt.Printf("Embedding length: %d\n", len(vector))
	//	fmt.Printf("First 5 values: %v\n", vector[:5])

	ctx.SetOutput(oResponse, vector)
	return true, nil
}
