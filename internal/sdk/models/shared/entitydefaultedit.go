// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type EntityDefaultEditViewType string

const (
	EntityDefaultEditViewTypeDefault EntityDefaultEditViewType = "default"
)

func (e EntityDefaultEditViewType) ToPointer() *EntityDefaultEditViewType {
	return &e
}
func (e *EntityDefaultEditViewType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "default":
		*e = EntityDefaultEditViewType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EntityDefaultEditViewType: %v", v)
	}
}

type EntityDefaultEdit struct {
	SearchParams map[string]string `json:"search_params,omitempty"`
	// List of attribute names that we show in the summary header
	SummaryAttributes []string                   `json:"summary_attributes,omitempty"`
	ViewType          *EntityDefaultEditViewType `json:"view_type,omitempty"`
}

func (o *EntityDefaultEdit) GetSearchParams() map[string]string {
	if o == nil {
		return nil
	}
	return o.SearchParams
}

func (o *EntityDefaultEdit) GetSummaryAttributes() []string {
	if o == nil {
		return nil
	}
	return o.SummaryAttributes
}

func (o *EntityDefaultEdit) GetViewType() *EntityDefaultEditViewType {
	if o == nil {
		return nil
	}
	return o.ViewType
}
