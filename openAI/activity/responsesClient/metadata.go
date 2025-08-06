package responsesClient

import (
	"github.com/project-flogo/core/data/coerce"
)

// Constants for identifying settings and inputs
const (
	sOutputFormat = "outputFormat"
	iAPIKey       = "apiKey"
	iModel        = "model"
	iPrompt       = "prompt"
	oResponse     = "response"
)

// Settings defines configuration options for your activity
type Settings struct {
	OutputFormat string `md:"OutputFormat"` // Flogo metadata tag
}

// FromMap populates the settings struct from a map.
func (s *Settings) FromMap(values map[string]interface{}) error {
	if values == nil {
		s.OutputFormat = "json"
		return nil
	}

	var err error

	if val, ok := values[sOutputFormat]; ok && val != nil {
		s.OutputFormat, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}
	if s.OutputFormat == "" {
		s.OutputFormat = "json"
	}

	return nil
}

// Input defines what data the activity receives
type Input struct {
	ApiKey map[string]interface{} `md:"apiKey, required"`
	Model  map[string]interface{} `md:"model, required"`
	Prompt map[string]interface{} `md:"prompt, required"`
}

// FromMap populates the struct from the activity's inputs.
func (i *Input) FromMap(values map[string]interface{}) error {

	if values == nil {
		return nil
	}

	// Todo Refactor this code to make efficient.
	var err error
	i.ApiKey, err = coerce.ToObject(values[iAPIKey])
	if err != nil {
		return err
	}

	i.Model, err = coerce.ToObject(values[iModel])
	if err != nil {
		return err
	}

	i.Prompt, err = coerce.ToObject(values[iPrompt])

	if err != nil {
		return err
	}

	return nil
}

// ToMap converts the struct to a map.
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		iAPIKey: i.ApiKey,
		iModel:  i.Model,
		iPrompt: i.Prompt,
	}
}

// Output defines what data the activity returns
type Output struct {
	Response string `md:"response"`
}

// ToMap converts the struct to a map.
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		oResponse: o.Response,
	}
}

// FromMap populates the struct from a map.
func (o *Output) FromMap(values map[string]interface{}) error {
	if values == nil {
		return nil
	}

	var err error
	if val, ok := values[oResponse]; ok && val != nil {
		o.Response, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}
	return nil
}
