// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"net/http"
)

type DeleteEntityRequest struct {
	// Activity to include in event feed
	ActivityID *string `queryParam:"style=form,explode=true,name=activity_id"`
	// Entity id
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Entity Type
	Slug string `pathParam:"style=simple,explode=false,name=slug"`
}

func (o *DeleteEntityRequest) GetActivityID() *string {
	if o == nil {
		return nil
	}
	return o.ActivityID
}

func (o *DeleteEntityRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DeleteEntityRequest) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

type DeleteEntityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteEntityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteEntityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteEntityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
