// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-taxonomy/internal/sdk/models/shared"
	"net/http"
)

type AttachActivityRequest struct {
	// Comma-separated list of entities which the activity primarily concerns
	Entities []string `queryParam:"style=form,explode=false,name=entities"`
	// Activity Id
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *AttachActivityRequest) GetEntities() []string {
	if o == nil {
		return nil
	}
	return o.Entities
}

func (o *AttachActivityRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type AttachActivityResponse struct {
	// Success
	BaseActivityItem *shared.BaseActivityItem
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *AttachActivityResponse) GetBaseActivityItem() *shared.BaseActivityItem {
	if o == nil {
		return nil
	}
	return o.BaseActivityItem
}

func (o *AttachActivityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *AttachActivityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *AttachActivityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
