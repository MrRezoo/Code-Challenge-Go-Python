package helpers

import "github.com/MrRezoo/code-challenge/api/validations"

type BaseHttpResponse struct {
	Result           any                            `json:"result"`
	Success          bool                           `json:"success"`
	Code             int                            `json:"code"`
	ValidationErrors *[]validations.ValidationError `json:"validation_errors"`
	Error            any                            `json:"error"`
}

func GenerateBaseResponse(result any, success bool, code int) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:  result,
		Success: success,
		Code:    code,
	}
}

func GenerateBaseResponseWithError(result any, success bool, code int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:  result,
		Success: success,
		Code:    code,
		Error:   err,
	}
}

func GenerateBaseResponseWithValidationErrors(result any, success bool, code int, err error) *BaseHttpResponse {
	return &BaseHttpResponse{
		Result:           result,
		Success:          success,
		Code:             code,
		ValidationErrors: validations.GetValidationErrors(err),
	}
}
