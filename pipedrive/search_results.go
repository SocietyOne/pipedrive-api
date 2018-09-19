package pipedrive

import (
	"context"
	"net/http"
)

// SearchResultsService handles search results related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/SearchResults
type SearchResultsService service

// SearchResult represents a Pipedrive search result.
type SearchResult struct {
	Type        string  `json:"type,omitempty"`
	ID          int     `json:"id,omitempty"`
	Source      string  `json:"source,omitempty"`
	ResultScore float64 `json:"result_score,omitempty"`
	Notes       struct {
		Count   int           `json:"count,omitempty"`
		Content []interface{} `json:"content,omitempty"`
	} `json:"notes,omitempty"`
	Fields struct {
		Count int           `json:"count,omitempty"`
		Names []interface{} `json:"names,omitempty"`
	} `json:"fields,omitempty"`
	Title   string      `json:"title,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

func (sr SearchResult) String() string {
	return Stringify(sr)
}

// SearchResults represents multiple search results response.
type SearchResults struct {
	Success        bool           `json:"success"`
	Data           []SearchResult `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

// SearchResultsListOptions specifices the optional parameters to the
// FieldType: dealField, personField, organizationField, productField
// ReturnItemIDs: 0, 1
// SearchResultsService.Search method.
type SearchResultsListOptions struct {
	Term       string `url:"term" json:"term,omitempty"`
	ItemType   string `url:"item_type" json:"item_type,omitempty"`
	Start      uint   `url:"start" json:"start,omitempty"`
	Limit      uint   `url:"limit" json:"limit,omitempty"`
	ExactMatch uint8  `url:"exact_match" json:"exact_match,omitempty"`

	FieldType     string `url:"field_type" json:"field_type,omitempty"`
	FieldKey      string `url:"field_key" json:"field_key,omitempty"`
	ReturnItemIDs int    `url:"return_item_ids" json:"return_item_ids,omitempty"`
}

// Search performs a search across the account and returns SearchResults.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/SearchResults/get_searchResults
func (s *SearchResultsService) Search(ctx context.Context, opt *SearchResultsListOptions) (*SearchResults, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/searchResults", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SearchResults

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// SearchByFieldValue performs a search on the value of a field
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/SearchResults/get_searchResults_field
func (s *SearchResultsService) SearchByFieldValue(ctx context.Context, opt *SearchResultsListOptions) (*SearchResults, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/searchResults/field", opt, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *SearchResults

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
