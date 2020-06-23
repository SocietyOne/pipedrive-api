package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

const (
	SearchItemFieldTypeDeal         SearchItemFieldType = "dealField"
	SearchItemFieldTypePerson       SearchItemFieldType = "personField"
	SearchItemFieldTypeOrganization SearchItemFieldType = "organizationField"
	SearchItemFieldTypeProduct      SearchItemFieldType = "productField"
)

type SearchItemFieldType string

// SearchItemFieldsOptions is used to configure a search request. Term, FieldType and FieldKey is required
type SearchItemFieldsOptions struct {
	Term          string              `url:"term"`                      // The search term to look for. Minimum 2 characters (or 1 if using exact_match).
	FieldType     SearchItemFieldType `url:"field_type,omitempty"`      // The type of the field to perform the search from
	ExactMatch    *bool               `url:"exact_match,omitempty"`     // When enabled, only full exact matches against the given term are returned. The search is case sensitive.
	FieldKey      string              `url:"field_key"`                 // The key of the field to search from. The field key can be obtained by fetching the list of the fields using any of the fields' API GET methods (dealFields, personFields, etc.).
	ReturnItemIDs *bool               `url:"return_item_ids,omitempty"` // Whether to return the IDs of the matching items or not. When not set or set to 0 or false, only distinct values of the searched field are returned. When set to 1 or true, the ID of each found item is returned.
	Start         *int                `url:"start,omitempty"`           // Pagination start.
	Limit         *int                `url:"limit,omitempty"`           // Items shown per page
}

// SearchItemFields searches an object by a specific field
// Pipedrive API https://developers.pipedrive.com/docs/api/v1/#!/ItemSearch/get_itemSearch_field
func (c *Client) SearchItemFields(ctx context.Context, opt SearchItemFieldsOptions, out ResponseModel) error {

	req, err := c.NewRequest(http.MethodGet, "/itemSearch/field", opt, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, out)
	if err != nil {
		return err
	}
	if !out.Successful() {
		return fmt.Errorf("not successful, error: %v", out.ErrorString())
	}

	return nil
}
