package methods

import (
	"encoding/json"
	"errors"
	"unsafe"

	"github.com/awnumar/memguard"

	"github.com/iotaledger/wasp-wallet-sdk/types"
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

func ParseResponseEnvelope(responseString []byte, responseErr error) (*ResponseEnvelope, error) {
	if responseErr != nil {
		return nil, responseErr
	}

	var responseEnvelope ResponseEnvelope
	if err := json.Unmarshal(responseString, &responseEnvelope); err != nil {
		return nil, err
	}

	if responseEnvelope.Type == "error" {
		var errorResponse JSONErrorResponse

		if err := json.Unmarshal(responseEnvelope.Payload, &errorResponse); err != nil {
			return nil, err
		}

		return nil, errors.New(errorResponse.ErrorMessage)
	}

	return &responseEnvelope, nil
}

// ParseResponse returns a typed response object
func ParseResponse[T any](responseString []byte, responseErr error) (*T, error) {
	responseEnvelope, err := ParseResponseEnvelope(responseString, responseErr)
	if err != nil {
		return nil, err
	}

	response := new(T)
	if err := json.Unmarshal(responseEnvelope.Payload, response); err != nil {
		return nil, err
	}

	return response, nil
}

// ParseResponseProtectedString handles the Payload as a []byte to mitigate unwanted string allocations
func ParseResponseProtectedString(responseString []byte, responseErr error) (*memguard.Enclave, error) {
	responseEnvelope, err := ParseResponseEnvelope(responseString, responseErr)
	if err != nil {
		return nil, err
	}

	var buffer string
	if err := json.Unmarshal(responseEnvelope.Payload, &buffer); err != nil {
		return nil, err
	}

	response := memguard.NewEnclave([]byte(buffer))
	bytes := *(*[]byte)(unsafe.Pointer(&buffer))
	memguard.WipeBytes(bytes)

	return response, nil
}

// ParseResponseStatus Returns true or false, whether the request succeeded or not.
func ParseResponseStatus(responseString []byte, responseErr error) (bool, error) {
	responseEnvelope, err := ParseResponseEnvelope(responseString, responseErr)
	if err != nil {
		return false, err
	}

	return responseEnvelope.Type == types.OperationSuccess, nil
}
