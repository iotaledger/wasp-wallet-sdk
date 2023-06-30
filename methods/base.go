package methods

import (
	"encoding/json"
	"errors"
)

type NoType any

type BaseRequest[T any] struct {
	Name string `json:"name" yaml:"name" mapstructure:"name"`
	Data T      `json:"data" yaml:"data" mapstructure:"data"`
}

func NewBaseRequest[T any](name string, data T) BaseRequest[T] {
	return BaseRequest[T]{
		Name: name,
		Data: data,
	}
}

func NewBaseRequestNoData(name string) BaseRequest[NoType] {
	return BaseRequest[NoType]{
		Name: name,
	}
}

type JSONErrorResponse struct {
	Type         string `json:"type"`
	ErrorMessage string `json:"error"`
}

type ResponseEnvelope struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func ParseResponse[T any](responseString string, responseErr error) (*T, error) {
	if responseErr != nil {
		return nil, responseErr
	}

	var responseEnvelope ResponseEnvelope

	if err := json.Unmarshal([]byte(responseString), &responseEnvelope); err != nil {
		return nil, err
	}

	if responseEnvelope.Type == "error" {
		var errorResponse JSONErrorResponse

		if err := json.Unmarshal(responseEnvelope.Payload, &errorResponse); err != nil {
			return nil, err
		}

		return nil, errors.New(errorResponse.ErrorMessage)
	}

	response := new(T)
	if err := json.Unmarshal(responseEnvelope.Payload, response); err != nil {
		return nil, err
	}

	return response, nil
}
