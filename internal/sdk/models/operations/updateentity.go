// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/internal/utils"
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type UpdateEntityRequest struct {
	Entity *shared.EntityInput `request:"mediaType=application/json"`
	// Activity to include in event feed
	ActivityID *string `queryParam:"style=form,explode=true,name=activity_id"`
	// Don't wait for updated entity to become available in Search API. Useful for large migrations
	Async *bool `default:"false" queryParam:"style=form,explode=true,name=async"`
	// Entity id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Entity Type
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

func (u UpdateEntityRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateEntityRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateEntityRequest) GetEntity() *shared.EntityInput {
	if o == nil {
		return nil
	}
	return o.Entity
}

func (o *UpdateEntityRequest) GetActivityID() *string {
	if o == nil {
		return nil
	}
	return o.ActivityID
}

func (o *UpdateEntityRequest) GetAsync() *bool {
	if o == nil {
		return nil
	}
	return o.Async
}

func (o *UpdateEntityRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateEntityRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

type UpdateEntityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Success
	EntityItem *shared.EntityItem
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateEntityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateEntityResponse) GetEntityItem() *shared.EntityItem {
	if o == nil {
		return nil
	}
	return o.EntityItem
}

func (o *UpdateEntityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateEntityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
