// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
)

// SelectAttributeConstraints - A set of constraints applicable to the attribute.
// These constraints should and will be enforced by the attribute renderer.
type SelectAttributeConstraints struct {
}

// SelectAttributeInfoHelpers - A set of configurations meant to document and assist the user in filling the attribute.
type SelectAttributeInfoHelpers struct {
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

func (o *SelectAttributeInfoHelpers) GetHintCustomComponent() *string {
	if o == nil {
		return nil
	}
	return o.HintCustomComponent
}

func (o *SelectAttributeInfoHelpers) GetHintText() *string {
	if o == nil {
		return nil
	}
	return o.HintText
}

func (o *SelectAttributeInfoHelpers) GetHintTextKey() *string {
	if o == nil {
		return nil
	}
	return o.HintTextKey
}

func (o *SelectAttributeInfoHelpers) GetHintTooltipPlacement() *string {
	if o == nil {
		return nil
	}
	return o.HintTooltipPlacement
}

type SelectAttribute1 struct {
	Title *string `json:"title,omitempty"`
	Value string  `json:"value"`
}

func (o *SelectAttribute1) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *SelectAttribute1) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

type SelectAttributeOptionsType string

const (
	SelectAttributeOptionsTypeSelectAttribute1 SelectAttributeOptionsType = "SelectAttribute_1"
	SelectAttributeOptionsTypeStr              SelectAttributeOptionsType = "str"
)

type SelectAttributeOptions struct {
	SelectAttribute1 *SelectAttribute1
	Str              *string

	Type SelectAttributeOptionsType
}

func CreateSelectAttributeOptionsSelectAttribute1(selectAttribute1 SelectAttribute1) SelectAttributeOptions {
	typ := SelectAttributeOptionsTypeSelectAttribute1

	return SelectAttributeOptions{
		SelectAttribute1: &selectAttribute1,
		Type:             typ,
	}
}

func CreateSelectAttributeOptionsStr(str string) SelectAttributeOptions {
	typ := SelectAttributeOptionsTypeStr

	return SelectAttributeOptions{
		Str:  &str,
		Type: typ,
	}
}

func (u *SelectAttributeOptions) UnmarshalJSON(data []byte) error {

	var selectAttribute1 SelectAttribute1 = SelectAttribute1{}
	if err := utils.UnmarshalJSON(data, &selectAttribute1, "", true, true); err == nil {
		u.SelectAttribute1 = &selectAttribute1
		u.Type = SelectAttributeOptionsTypeSelectAttribute1
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = SelectAttributeOptionsTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SelectAttributeOptions", string(data))
}

func (u SelectAttributeOptions) MarshalJSON() ([]byte, error) {
	if u.SelectAttribute1 != nil {
		return utils.MarshalJSON(u.SelectAttribute1, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type SelectAttributeOptions: all fields are null")
}

type SelectAttributeType string

const (
	SelectAttributeTypeSelect SelectAttributeType = "select"
	SelectAttributeTypeRadio  SelectAttributeType = "radio"
)

func (e SelectAttributeType) ToPointer() *SelectAttributeType {
	return &e
}
func (e *SelectAttributeType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "select":
		fallthrough
	case "radio":
		*e = SelectAttributeType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SelectAttributeType: %v", v)
	}
}

// SelectAttribute - Dropdown select
type SelectAttribute struct {
	Purpose []string `json:"_purpose,omitempty"`
	// Allow arbitrary input values in addition to provided options
	AllowAny *bool `json:"allow_any,omitempty"`
	// A set of constraints applicable to the attribute.
	// These constraints should and will be enforced by the attribute renderer.
	//
	Constraints  *SelectAttributeConstraints `json:"constraints,omitempty"`
	DefaultValue any                         `json:"default_value,omitempty"`
	Deprecated   *bool                       `default:"false" json:"deprecated"`
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
	InfoHelpers *SelectAttributeInfoHelpers `json:"info_helpers,omitempty"`
	Label       string                      `json:"label"`
	Layout      *string                     `json:"layout,omitempty"`
	Name        string                      `json:"name"`
	Options     []SelectAttributeOptions    `json:"options,omitempty"`
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
	Type           *SelectAttributeType `json:"type,omitempty"`
	ValueFormatter *string              `json:"value_formatter,omitempty"`
}

func (s SelectAttribute) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SelectAttribute) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SelectAttribute) GetPurpose() []string {
	if o == nil {
		return nil
	}
	return o.Purpose
}

func (o *SelectAttribute) GetAllowAny() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAny
}

func (o *SelectAttribute) GetConstraints() *SelectAttributeConstraints {
	if o == nil {
		return nil
	}
	return o.Constraints
}

func (o *SelectAttribute) GetDefaultValue() any {
	if o == nil {
		return nil
	}
	return o.DefaultValue
}

func (o *SelectAttribute) GetDeprecated() *bool {
	if o == nil {
		return nil
	}
	return o.Deprecated
}

func (o *SelectAttribute) GetEntityBuilderDisableEdit() *bool {
	if o == nil {
		return nil
	}
	return o.EntityBuilderDisableEdit
}

func (o *SelectAttribute) GetFeatureFlag() *string {
	if o == nil {
		return nil
	}
	return o.FeatureFlag
}

func (o *SelectAttribute) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *SelectAttribute) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *SelectAttribute) GetHideLabel() *bool {
	if o == nil {
		return nil
	}
	return o.HideLabel
}

func (o *SelectAttribute) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *SelectAttribute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SelectAttribute) GetInfoHelpers() *SelectAttributeInfoHelpers {
	if o == nil {
		return nil
	}
	return o.InfoHelpers
}

func (o *SelectAttribute) GetLabel() string {
	if o == nil {
		return ""
	}
	return o.Label
}

func (o *SelectAttribute) GetLayout() *string {
	if o == nil {
		return nil
	}
	return o.Layout
}

func (o *SelectAttribute) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SelectAttribute) GetOptions() []SelectAttributeOptions {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *SelectAttribute) GetOrder() *int64 {
	if o == nil {
		return nil
	}
	return o.Order
}

func (o *SelectAttribute) GetPlaceholder() *string {
	if o == nil {
		return nil
	}
	return o.Placeholder
}

func (o *SelectAttribute) GetPreviewValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.PreviewValueFormatter
}

func (o *SelectAttribute) GetProtected() *bool {
	if o == nil {
		return nil
	}
	return o.Protected
}

func (o *SelectAttribute) GetReadonly() *bool {
	if o == nil {
		return nil
	}
	return o.Readonly
}

func (o *SelectAttribute) GetRenderCondition() *string {
	if o == nil {
		return nil
	}
	return o.RenderCondition
}

func (o *SelectAttribute) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *SelectAttribute) GetSettingsFlag() []SettingFlag {
	if o == nil {
		return nil
	}
	return o.SettingsFlag
}

func (o *SelectAttribute) GetShowInTable() *bool {
	if o == nil {
		return nil
	}
	return o.ShowInTable
}

func (o *SelectAttribute) GetSortable() *bool {
	if o == nil {
		return nil
	}
	return o.Sortable
}

func (o *SelectAttribute) GetType() *SelectAttributeType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *SelectAttribute) GetValueFormatter() *string {
	if o == nil {
		return nil
	}
	return o.ValueFormatter
}
