package promtengineclient

import (
	"encoding/json"

	model "github.com/BinaryLogos/http-models"
)

// Generate sends a request to the promt engine /generate/promt endpoint
func (c *client) Generate(req model.PromptRequest) (*model.PromptResponse, error) {
	request, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := c.post(generateEndpoint, string(request))
	if err != nil {
		return nil, err
	}

	var resp model.PromptResponse

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
