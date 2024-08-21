// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
)

// ConsentAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type ConsentAttributeConstraints struct {
}

// ConsentAttributeInfoHelpers - A set of configurations meant to document and assist the user in filling the attribute.
type ConsentAttributeInfoHelpers struct {
	// The name of the custom component to be used as the hint helper.
	// The component should be registered in the `@epilot360/entity-ui` on the index of the components directory.
	// When specified it overrides the `hint_text` or `hint_text_key` configuration.
	//
	HintCustomComponent *string `json:"hint_custom_component,omitempty"`
	// The text to be displayed in the attribute hint helper.
	// When specified it overrides the `hint_text_key` configuration.
	//
	HintText *string `json:"hint_text,omitempty"`
	// The key of the hint text to be displayed in the attribute hint helper.
	// The key should be a valid i18n key.
	//
	HintTextKey *string `json:"hint_text_key,omitempty"`
	// The placement of the hint tooltip.
	// The value should be a valid `@mui/core` tooltip placement.
	//
	HintTooltipPlacement *string `json:"hint_tooltip_placement,omitempty"`
}

func (o *ConsentAttributeInfoHelpers) GetHintCustomComponent() *string {
	if o == nil {
		return nil
	}
	return o.HintCustomComponent
}

func (o *ConsentAttributeInfoHelpers) GetHintText() *string {
	if o == nil {
		return nil
	}
	return o.HintText
}

func (o *ConsentAttributeInfoHelpers) GetHintTextKey() *string {
	if o == nil {
		return nil
	}
	return o.HintTextKey
}

func (o *ConsentAttributeInfoHelpers) GetHintTooltipPlacement() *string {
	if o == nil {
		return nil
	}
	return o.HintTooltipPlacement
}

type ConsentAttributeType string

const (
	ConsentAttributeTypeConsent ConsentAttributeType = "consent"
)

func (e ConsentAttributeType) ToPointer() *ConsentAttributeType {
	return &e
}
func (e *ConsentAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consent":
		*e = ConsentAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ConsentAttributeType: %v", v)
	}
}

// ConsentAttribute - Consent Management
type ConsentAttribute struct {
	Purpose []string `json:"_purpose,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *ConsentAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue any                          `json:"default_value,omitempty"`
	Deprecated   *bool                        `default:"false" json:"deprecated"`
	// Setting to `true` disables editing the attribute on the entity builder UI
	EntityBuilderDisableEdit *bool `default:"false" json:"entity_builder_disable_edit"`
	// This attribute should only be active when the feature flag is enabled
	FeatureFlag *string `json:"feature_flag,omitempty"`
	// Which group the attribute should appear in. Accepts group ID or group name
	Group *string `json:"group,omitempty"`
	// Do not render attribute in entity views
	Hidden *bool `default:"false" json:"hidden"`
	// When set to true, will hide the label of the field.
	HideLabel *bool `json:"hide_label,omitempty"`
	// Code name of the icon to used to represent this attribute.
	// The value must be a valid @epilot/base-elements Icon name
	//
	Icon *string `json:"icon,omitempty"`
	// ID for the entity attribute
	ID          *string  `json:"id,omitempty"`
	Identifiers []string `json:"identifiers,omitempty"`
	// A set of configurations meant to document and assist the user in filling the attribute.
	InfoHelpers *ConsentAttributeInfoHelpers `json:"info_helpers,omitempty"`
	Label       string                       `json:"label"`
	Layout      *string                      `json:"layout,omitempty"`
	Name        string                       `json:"name"`
	// Attribute sort order (ascending) in group
	Order                 *int64  `json:"order,omitempty"`
	Placeholder           *string `json:"placeholder,omitempty"`
	PreviewValueFormatter *string `json:"preview_value_formatter,omitempty"`
	// Setting to `true` prevents the attribute from being modified / deleted
	Protected *bool `json:"protected,omitempty"`
	Readonly  *bool `default:"false" json:"readonly"`
	// Defines the conditional rendering expression for showing this field.
	// When a valid expression is parsed, their evaluation defines the visibility of this attribute.
	// Note: Empty or invalid expression have no effect on the field visibility.
	//
	RenderCondition *string `json:"render_condition,omitempty"`
	Required        *bool   `default:"false" json:"required"`
	// This attribute should only be active when all the settings have the correct value
	SettingsFlag []SettingFlag `json:"settings_flag,omitempty"`
	// Render as a column in table views. When defined, overrides `hidden`
	ShowInTable *bool `json:"show_in_table,omitempty"`
	// Allow sorting by this attribute in table views if `show_in_table` is true
	Sortable       *bool                `default:"true" json:"sortable"`
	Topic          string               `json:"topic"`
	Type           ConsentAttributeType `json:"type"`
	ValueFormatter *string              `json:"value_formatter,omitempty"`
}

func (c ConsentAttribute) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *ConsentAttribute) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *ConsentAttribute) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *ConsentAttribute) GetConstraints() *ConsentAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *ConsentAttribute) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *ConsentAttribute) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *ConsentAttribute) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *ConsentAttribute) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *ConsentAttribute) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *ConsentAttribute) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *ConsentAttribute) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *ConsentAttribute) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *ConsentAttribute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *ConsentAttribute) GetIdentifiers() []string {
	if o == nil {
		return nil
	}
	return o.Identifiers
}

func (o *ConsentAttribute) GetInfoHelpers() *ConsentAttributeInfoHelpers {
	if o == nil {
		return nil
	}
	return o.InfoHelpers
}

func (o *ConsentAttribute) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *ConsentAttribute) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *ConsentAttribute) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *ConsentAttribute) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *ConsentAttribute) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *ConsentAttribute) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *ConsentAttribute) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *ConsentAttribute) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *ConsentAttribute) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *ConsentAttribute) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *ConsentAttribute) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *ConsentAttribute) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *ConsentAttribute) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *ConsentAttribute) GetTopic() string {
	if o == nil {
		return ""
	}
	return o.Topic
}

func (o *ConsentAttribute) GetType() ConsentAttributeType {
	if o == nil {
		return ConsentAttributeType("")
	}
	return o.Type
}

func (o *ConsentAttribute) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}
