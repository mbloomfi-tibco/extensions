package chatGPT

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting,required"`
}

type Input struct {
	apiKey string `md:"apiKey,required"`
	model  string `md:"model,required"`
	prompt string `md:"prompt,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strApiKey, _ := coerce.ToString(values["apiKey"])
	r.apiKey = strApiKey
	strModel, _ := coerce.ToString(values["model"])
	r.apiKey = strModel
	strPrompt, _ := coerce.ToString(values["prompt"])
	r.apiKey = strPrompt
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"apiKey": r.apiKey,
		"model":  r.model,
		"prompt": r.prompt,
	}
}

type Output struct {
	response string `md:"response"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["response"])
	o.response = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"response": o.response,
	}
}
