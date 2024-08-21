// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
	"time"
)

type EntityItem struct {
	AdditionalProperties any `additionalProperties:"true" json:"-"`
	// Access control list (ACL) for an entity. Defines sharing access to external orgs or users.
	ACL       *EntityACL `json:"_acl,omitempty"`
	CreatedAt *time.Time `json:"_created_at"`
	ID        string     `json:"_id"`
	// Organization Id the entity belongs to
	Org    string        `json:"_org"`
	Owners []EntityOwner `json:"_owners,omitempty"`
	// URL-friendly identifier for the entity schema
	Schema string   `json:"_schema"`
	Tags   []string `json:"_tags,omitempty"`
	// Title of entity
	Title     *string    `json:"_title"`
	UpdatedAt *time.Time `json:"_updated_at"`
}

func (e EntityItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EntityItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EntityItem) GetAdditionalProperties() any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *EntityItem) GetACL() *EntityACL {
	if o == nil {
		return nil
	}
	return o.ACL
}

func (o *EntityItem) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *EntityItem) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *EntityItem) GetOrg() string {
	if o == nil {
		return ""
	}
	return o.Org
}

func (o *EntityItem) GetOwners() []EntityOwner {
	if o == nil {
		return nil
	}
	return o.Owners
}

func (o *EntityItem) GetSchema() string {
	if o == nil {
		return ""
	}
	return o.Schema
}

func (o *EntityItem) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *EntityItem) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *EntityItem) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
