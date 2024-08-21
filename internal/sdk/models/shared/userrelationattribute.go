// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
)

// UserRelationAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type UserRelationAttributeConstraints struct {
}

// UserRelationAttributeInfoHelpers - A set of configurations meant to document and assist the user in filling the attribute.
type UserRelationAttributeInfoHelpers struct {
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

func (o *UserRelationAttributeInfoHelpers) GetHintCustomComponent() *string {
	if o == nil {
		return nil
	}
	return o.HintCustomComponent
}

func (o *UserRelationAttributeInfoHelpers) GetHintText() *string {
	if o == nil {
		return nil
	}
	return o.HintText
}

func (o *UserRelationAttributeInfoHelpers) GetHintTextKey() *string {
	if o == nil {
		return nil
	}
	return o.HintTextKey
}

func (o *UserRelationAttributeInfoHelpers) GetHintTooltipPlacement() *string {
	if o == nil {
		return nil
	}
	return o.HintTooltipPlacement
}

type UserRelationAttributeType string

const (
	UserRelationAttributeTypeRelationUser UserRelationAttributeType = "relation_user"
)

func (e UserRelationAttributeType) ToPointer() *UserRelationAttributeType {
	return &e
}
func (e *UserRelationAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "relation_user":
		*e = UserRelationAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UserRelationAttributeType: %v", v)
	}
}

// UserRelationAttribute - User Relationship
type UserRelationAttribute struct {
	Purpose []string `json:"_purpose,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *UserRelationAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue any                               `json:"default_value,omitempty"`
	Deprecated   *bool                             `default:"false" json:"deprecated"`
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
	ID *string `json:"id,omitempty"`
	// A set of configurations meant to document and assist the user in filling the attribute.
	InfoHelpers *UserRelationAttributeInfoHelpers `json:"info_helpers,omitempty"`
	Label       string                            `json:"label"`
	Layout      *string                           `json:"layout,omitempty"`
	Multiple    *bool                             `default:"false" json:"multiple"`
	Name        string                            `json:"name"`
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
	Sortable       *bool                      `default:"true" json:"sortable"`
	Type           *UserRelationAttributeType `json:"type,omitempty"`
	ValueFormatter *string                    `json:"value_formatter,omitempty"`
}

func (u UserRelationAttribute) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UserRelationAttribute) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *UserRelationAttribute) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *UserRelationAttribute) GetConstraints() *UserRelationAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *UserRelationAttribute) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *UserRelationAttribute) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *UserRelationAttribute) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *UserRelationAttribute) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *UserRelationAttribute) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *UserRelationAttribute) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *UserRelationAttribute) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *UserRelationAttribute) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *UserRelationAttribute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UserRelationAttribute) GetInfoHelpers() *UserRelationAttributeInfoHelpers {
	if o == nil {
		return nil
	}
	return o.InfoHelpers
}

func (o *UserRelationAttribute) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *UserRelationAttribute) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *UserRelationAttribute) GetMultiple() *bool {
	if o == nil {
		return nil
	}
	return o.Multiple
}

func (o *UserRelationAttribute) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UserRelationAttribute) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *UserRelationAttribute) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *UserRelationAttribute) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *UserRelationAttribute) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *UserRelationAttribute) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *UserRelationAttribute) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *UserRelationAttribute) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *UserRelationAttribute) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *UserRelationAttribute) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *UserRelationAttribute) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *UserRelationAttribute) GetType() *UserRelationAttributeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UserRelationAttribute) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}
