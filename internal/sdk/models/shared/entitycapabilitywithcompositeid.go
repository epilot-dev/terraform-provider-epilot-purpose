// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
)

// EntityCapabilityWithCompositeIDRequiredPermission - Require a permission to display UI hook
type EntityCapabilityWithCompositeIDRequiredPermission struct {
	Action   string  `json:"action"`
	Resource *string `json:"resource,omitempty"`
}

func (o *EntityCapabilityWithCompositeIDRequiredPermission) GetAction() string {
	if o == nil {
		return ""
	}
	return o.Action
}

func (o *EntityCapabilityWithCompositeIDRequiredPermission) GetResource() *string {
	if o == nil {
		return nil
	}
	return o.Resource
}

type EntityCapabilityWithCompositeIDUIHooks struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// the component to be dynamically loaded
	Component *string `json:"component,omitempty"`
	// Whether capability should be disabled
	Disabled *bool `json:"disabled,omitempty"`
	// Sets the group expand/collapse default state
	GroupExpanded *bool `json:"group_expanded,omitempty"`
	// Specific to Activity pilot
	Header *bool `json:"header,omitempty"`
	// name of the hook to use
	Hook string `json:"hook"`
	// Preview icon name(As in Base elements) for the capability
	Icon *string `json:"icon,omitempty"`
	// package to be imported
	Import *string `json:"import,omitempty"`
	// render order (ascending)
	Order           *int64  `json:"order,omitempty"`
	RenderCondition *string `json:"render_condition,omitempty"`
	// Require a permission to display UI hook
	RequiredPermission *EntityCapabilityWithCompositeIDRequiredPermission `json:"requiredPermission,omitempty"`
	// route for specified capability
	Route *string `json:"route,omitempty"`
	Title *string `json:"title,omitempty"`
}

func (e EntityCapabilityWithCompositeIDUIHooks) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EntityCapabilityWithCompositeIDUIHooks) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetComponent() *string {
	if o == nil {
		return nil
	}
	return o.Component
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetDisabled() *bool {
	if o == nil {
		return nil
	}
	return o.Disabled
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetGroupExpanded() *bool {
	if o == nil {
		return nil
	}
	return o.GroupExpanded
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetHeader() *bool {
	if o == nil {
		return nil
	}
	return o.Header
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetHook() string {
	if o == nil {
		return ""
	}
	return o.Hook
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetImport() *string {
	if o == nil {
		return nil
	}
	return o.Import
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetRequiredPermission() *EntityCapabilityWithCompositeIDRequiredPermission {
	if o == nil {
		return nil
	}
	return o.RequiredPermission
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetRoute() *string {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *EntityCapabilityWithCompositeIDUIHooks) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

// EntityCapabilityWithCompositeID - a readonly computed ID for the entity capability including schema slug and the capability ID
type EntityCapabilityWithCompositeID struct {
	Purpose     []string    `json:"_purpose,omitempty"`
	Attributes  []Attribute `json:"attributes,omitempty"`
	CompositeID *string     `json:"composite_id,omitempty"`
	// This capability should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// ID for the entity capability
	ID *string `json:"id,omitempty"`
	// Unique name for the capability
	Name string `json:"name"`
	// This capability should only be active when all the settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
	// Human readable title of the capability
	Title   *string                                  `json:"title,omitempty"`
	UIHooks []EntityCapabilityWithCompositeIDUIHooks `json:"ui_hooks,omitempty"`
}

func (o *EntityCapabilityWithCompositeID) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *EntityCapabilityWithCompositeID) GetAttributes() []Attribute {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *EntityCapabilityWithCompositeID) GetCompositeID() *string {
	if o == nil {
		return nil
	}
	return o.CompositeID
}

func (o *EntityCapabilityWithCompositeID) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *EntityCapabilityWithCompositeID) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EntityCapabilityWithCompositeID) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *EntityCapabilityWithCompositeID) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *EntityCapabilityWithCompositeID) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *EntityCapabilityWithCompositeID) GetUIHooks() []EntityCapabilityWithCompositeIDUIHooks {
	if o == nil {
		return nil
	}
	return o.UIHooks
}

// EntityCapabilityWithCompositeIDInput - a readonly computed ID for the entity capability including schema slug and the capability ID
type EntityCapabilityWithCompositeIDInput struct {
	Purpose    []string         `json:"_purpose,omitempty"`
	Attributes []AttributeInput `json:"attributes,omitempty"`
	// This capability should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// ID for the entity capability
	ID *string `json:"id,omitempty"`
	// Unique name for the capability
	Name string `json:"name"`
	// This capability should only be active when all the settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
	// Human readable title of the capability
	Title   *string                                  `json:"title,omitempty"`
	UIHooks []EntityCapabilityWithCompositeIDUIHooks `json:"ui_hooks,omitempty"`
}

func (o *EntityCapabilityWithCompositeIDInput) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *EntityCapabilityWithCompositeIDInput) GetAttributes() []AttributeInput {
	if o == nil {
		return nil
	}
	return o.Attributes
}

func (o *EntityCapabilityWithCompositeIDInput) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *EntityCapabilityWithCompositeIDInput) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EntityCapabilityWithCompositeIDInput) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *EntityCapabilityWithCompositeIDInput) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *EntityCapabilityWithCompositeIDInput) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *EntityCapabilityWithCompositeIDInput) GetUIHooks() []EntityCapabilityWithCompositeIDUIHooks {
	if o == nil {
		return nil
	}
	return o.UIHooks
}
