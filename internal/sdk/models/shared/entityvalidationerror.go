// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

// EntityValidationErrorParams - Additional parameters for the error
type EntityValidationErrorParams struct {
	// The type of the error
	Type *string `json:"type,omitempty"`
}

func (o *EntityValidationErrorParams) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

// EntityValidationError - Validation error for an entity attribute
type EntityValidationError struct {
	// Error code identifier
	Code string `json:"code"`
	// A human-readable message describing the error
	Message string `json:"message"`
	// Additional parameters for the error
	Params EntityValidationErrorParams `json:"params"`
	// The path to the attribute that failed validation
	Path []string `json:"path"`
}

func (o *EntityValidationError) GetCode() string {
	if o == nil {
		return ""
	}
	return o.Code
}

func (o *EntityValidationError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *EntityValidationError) GetParams() EntityValidationErrorParams {
	if o == nil {
		return EntityValidationErrorParams{}
	}
	return o.Params
}

func (o *EntityValidationError) GetPath() []string {
	if o == nil {
		return []string{}
	}
	return o.Path
}
