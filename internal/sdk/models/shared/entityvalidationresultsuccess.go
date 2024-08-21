// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type EntityValidationResultSuccessStatus string

const (
	EntityValidationResultSuccessStatusSuccess EntityValidationResultSuccessStatus = "success"
)

func (e EntityValidationResultSuccessStatus) ToPointer() *EntityValidationResultSuccessStatus {
	return &e
}
func (e *EntityValidationResultSuccessStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		*e = EntityValidationResultSuccessStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EntityValidationResultSuccessStatus: %v", v)
	}
}

// EntityValidationResultSuccess - Validation result for a successful validation
type EntityValidationResultSuccess struct {
	Errors []EntityValidationError             `json:"errors"`
	Status EntityValidationResultSuccessStatus `json:"status"`
}

func (o *EntityValidationResultSuccess) GetErrors() []EntityValidationError {
	if o == nil {
		return []EntityValidationError{}
	}
	return o.Errors
}

func (o *EntityValidationResultSuccess) GetStatus() EntityValidationResultSuccessStatus {
	if o == nil {
		return EntityValidationResultSuccessStatus("")
	}
	return o.Status
}
