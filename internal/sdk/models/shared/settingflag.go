// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type SettingFlag struct {
	// Whether the setting should be enabled or not
	Enabled *bool `json:"enabled,omitempty"`
	// The name of the organization setting to check
	Name *string `json:"name,omitempty"`
}

func (o *SettingFlag) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *SettingFlag) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}
