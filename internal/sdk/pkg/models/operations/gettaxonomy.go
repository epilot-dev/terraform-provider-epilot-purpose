// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/epilot-dev/terraform-provider-epilot-purpose/internal/sdk/pkg/models/shared"
	"net/http"
)

type GetTaxonomyRequest struct {
	// Taxonomy slug
	TaxonomySlug string `pathParam:"style=simple,explode=false,name=taxonomySlug"`
}

func (o *GetTaxonomyRequest) GetTaxonomySlug() string {
	if o == nil {
		return ""
	}
	return o.TaxonomySlug
}

type GetTaxonomyResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Taxonomy
	Taxonomy *shared.Taxonomy
}

func (o *GetTaxonomyResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTaxonomyResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTaxonomyResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTaxonomyResponse) GetTaxonomy() *shared.Taxonomy {
	if o == nil {
		return nil
	}
	return o.Taxonomy
}
