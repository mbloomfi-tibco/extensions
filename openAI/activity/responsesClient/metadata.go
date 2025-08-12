package responsesClient

/*
* Copyright © 2023 - 2024. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"github.com/project-flogo/core/data/coerce"
)

// Constants for identifying settings and inputs
const (
	sAPIKey       = "apiKey"
	sInputFormat  = "inputFormat"
	sOutputFormat = "outputFormat"
	iModel        = "model"
	iPrompt       = "prompt"
	iTool         = "tool"
	iBase64String = "base64String"
	oResponse     = "response"
)

// Settings defines configuration options for your activity
type Settings struct {
	ApiKey       string `md:"apiKey, required"`
	InputFormat  string `md:"inputFormat"`  // Flogo metadata tag
	OutputFormat string `md:"outputFormat"` // Flogo metadata tag
}

// FromMap populates the settings struct from a map.
func (s *Settings) FromMap(values map[string]interface{}) error {
	if values == nil {
		s.ApiKey = ""
		s.InputFormat = "text"
		s.OutputFormat = "json"
		return nil
	}

	var err error

	if val, ok := values[sInputFormat]; ok && val != nil {
		s.InputFormat, err = coerce.ToString(val)

		if err != nil {
			return err
		}

		if s.InputFormat == "" {
			s.InputFormat = "text"
		}
	}

	if val, ok := values[sOutputFormat]; ok && val != nil {
		s.OutputFormat, err = coerce.ToString(val)

		if err != nil {
			return err
		}

		if s.OutputFormat == "" {
			s.OutputFormat = "json"
		}
	}

	s.ApiKey, err = coerce.ToString(values[sAPIKey])
	if err != nil {
		return err
	}

	return nil
}

// Input defines what data the activity receives
type Input struct {
	Model        map[string]interface{} `md:"model, required"`
	Prompt       map[string]interface{} `md:"prompt, required"`
	Tool         map[string]interface{} `md:"tool, required"`
	Base64String map[string]interface{} `md:"base64String, required"`
}

// FromMap populates the struct from the activity's inputs.
func (i *Input) FromMap(values map[string]interface{}) error {

	if values == nil {
		return nil
	}

	// Todo Refactor this code to make efficient.
	var err error

	i.Model, err = coerce.ToObject(values[iModel])
	if err != nil {
		return err
	}

	i.Prompt, err = coerce.ToObject(values[iPrompt])
	if err != nil {
		return err
	}

	i.Tool, err = coerce.ToObject(values[iTool])
	if err != nil {
		return err
	}

	i.Base64String, err = coerce.ToObject(values[iBase64String])
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

// ToMap converts the struct to a map.
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		iModel:        i.Model,
		iPrompt:       i.Prompt,
		iTool:         i.Tool,
		iBase64String: i.Base64String,
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
