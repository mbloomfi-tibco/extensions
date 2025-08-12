package imagesClient

/*
* Copyright © 2023 - 2024. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"os"

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
	inputFormat  string
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
	prompt := ctx.GetInput(iPrompt).(string)
	fileName := ctx.GetInput(iFileName).(string)
	//tool := ctx.GetInput(iTool).(string)

	if a.apiKey == "" {
		log.Fatal("Missing openAPI key")
	}

	oaiClient := openai.NewClient(
		option.WithAPIKey(a.apiKey),
	)

	clientCtx := context.Background()

	// Request an image
	imgResp, err := oaiClient.Images.Generate(clientCtx, openai.ImageGenerateParams{
		Model:  openai.ImageModelGPTImage1, // "gpt-image-1"
		Prompt: prompt,
		Size:   "1024x1024",
		// ResponseFormat: openai.F("b64_json"), // Optional, defaults to base64
	})
	if err != nil {
		log.Printf("Image generation error: %v\n", err)
		return false, err
	}

	if len(imgResp.Data) == 0 {
		log.Printf("No image data returned")
		return false, fmt.Errorf("no image data returned from OpenAI API")
	}

	// Get the Base64 data
	b64Data := imgResp.Data[0].B64JSON

	// Decode the Base64 string to bytes
	imgBytes, err := base64.StdEncoding.DecodeString(b64Data)
	if err != nil {
		log.Printf("Base64 decode error: %v\n", err)
		return false, err
	}

	// Save to file

	if err := os.WriteFile(fileName, imgBytes, 0644); err != nil {
		log.Printf("File save error: %v\n", err)
		return true, err
	}

	log.Printf("Image saved as %s\n", fileName)

	ctx.SetOutput(oResponse, "Image generated successfully")
	return true, nil
}
