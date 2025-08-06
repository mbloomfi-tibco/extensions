package chatGPT

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
	ASetting string `md:"aSetting,required"`
}

// // Input corresponds to activity.json inputs
// type Input struct {
// 	apiKey string `md:"apiKey,required"`
// 	model  string `md:"model,required"`
// 	prompt string `md:"prompt,required"`
// }

// // FromMap converts a map to Input struct
// func (i *Input) FromMap(values map[string]interface{}) error {
// 	var err error

// 	i.apiKey, err = coerce.ToString(values["apiKey"])
// 	if err != nil {
// 		return err
// 	}

// 	i.model, err = coerce.ToString(values["model"])
// 	if err != nil {
// 		return err
// 	}

// 	i.prompt, err = coerce.ToString(values["prompt"])
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (r *Input) ToMap() map[string]interface{} {
// 	return map[string]interface{}{
// 		"apiKey": r.apiKey,
// 		"model":  r.model,
// 		"prompt": r.prompt,
// 	}
// }

type Input struct {
	AnInput string `md:"anInput,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["anInput"])
	r.AnInput = strVal
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"anInput": r.AnInput,
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
