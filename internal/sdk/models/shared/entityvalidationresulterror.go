// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Status string

const (
	StatusError Status = "error"
)

func (e Status) ToPointer() *Status {
	return &e
}
func (e *Status) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "error":
		*e = Status(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Status: %v", v)
	}
}

// EntityValidationResultError - Validation result for a failed validation
type EntityValidationResultError struct {
	Errors []EntityValidationError `json:"errors"`
	Status Status                  `json:"status"`
}

func (o *EntityValidationResultError) GetErrors() []EntityValidationError {
	if o == nil {
		return []EntityValidationError{}
	}
	return o.Errors
}

func (o *EntityValidationResultError) GetStatus() Status {
	if o == nil {
		return Status("")
	}
	return o.Status
}
