package chatGPT

import "github.com/project-flogo/core/activity"

var jsonMetadata = `{
  "name": "chatgpt-client",
  "version": "1.0.0",
  "title": "ChatGPT Client",
  "description": "Flogo activity to call ChatGPT API",
  "inputs":[
    {
      "name": "apiKey",
      "type": "string"
    },
    {
      "name": "model",
      "type": "string"
    },
    {
      "name": "prompt",
      "type": "string"
    }
  ],
  "outputs":[
    {
      "name": "response",
      "type": "string"
    }
  ]
}`

// NewMetadata returns the activity metadata
func NewMetadata() *activity.Metadata {
	return activity.NewMetadataFromJson(jsonMetadata)
}
