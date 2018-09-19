package integration

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SocietyOne/pipedrive-api/pipedrive"
	"github.com/stretchr/testify/assert"
)

func TestSearchResults_Search(t *testing.T) {

	responseSwitch := 0
	testApiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch responseSwitch {
		case 0:
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":[{"type":"deal","id":33484,"source":"elastic2","result_score":3.3774848,"notes":{},"fields":{},"title":"New AppForm Test","details":{"currency":"AUD","org_address":"","org_id":null,"org_name":null,"person_id":"27751","person_name":"Tom New APITest","stage_id":15,"stage_name":"New Application","status":"open","value":54321,"value_formatted":"$54,321"}},{"type":"deal","id":33489,"source":"elastic2","result_score":3.3774848,"notes":{},"fields":{},"title":"New AppForm Test","details":{"currency":"AUD","org_address":"","org_id":null,"org_name":null,"person_id":"27751","person_name":"Tom New APITest","stage_id":15,"stage_name":"New Application","status":"open","value":543210,"value_formatted":"$543,210"}},{"type":"deal","id":33129,"source":"elastic2","result_score":3.3336477,"notes":{},"fields":{},"title":"New AppForm Test","details":{"currency":"AUD","org_address":"","org_id":null,"org_name":null,"person_id":"27751","person_name":"Tom New APITest","stage_id":6,"stage_name":"Call New Quotes Submitted","status":"open","value":4321,"value_formatted":"$4,321"}},{"type":"deal","id":33147,"source":"elastic2","result_score":3.3336477,"notes":{},"fields":{},"title":"New AppForm Test","details":{"currency":"AUD","org_address":"","org_id":null,"org_name":null,"person_id":"27751","person_name":"Tom New APITest","stage_id":15,"stage_name":"New Application","status":"open","value":54321,"value_formatted":"$54,321"}},{"type":"deal","id":33148,"source":"elastic2","result_score":3.3336477,"notes":{},"fields":{},"title":"New AppForm Test","details":{"currency":"AUD","org_address":"","org_id":null,"org_name":null,"person_id":"27751","person_name":"Tom New APITest","stage_id":15,"stage_name":"New Application","status":"open","value":54321,"value_formatted":"$54,321"}}],"additional_data":{"user":{"profile":{"id":0,"email":"","name":"","is_admin":false,"default_currency":"","icon_url":null,"activated":false},"locale":{"language":"","country":"","uses_12_hour_clock":false},"timezone":{"name":"","offset":0}},"multiple_companies":false,"default_company_id":0,"company_id":0,"since_timestamp":"","last_timestamp_on_page":"","pagination":{"start":0,"limit":0,"more_items_in_collection":false}}}`))
		default:
			t.Error("Too many requests")
		}
		responseSwitch = responseSwitch + 1
	}))

	config := &pipedrive.Config{
		APIKey:  "1",
		BaseURL: testApiServer.URL,
	}
	mockClient := pipedrive.NewClient(config)

	result, _, err := mockClient.SearchResults.Search(context.Background(), &pipedrive.SearchResultsListOptions{
		Term: "AppForm",
	})

	if err != nil {
		t.Errorf("Could not get search results: %v", err)
	}

	if result.Success != true {
		t.Error("Got invalid search results")
	}

	for _, match := range result.Data {
		byt, err := json.Marshal(match)
		if err != nil {
			t.Error("Marshal error")
		}
		assert.Contains(t, string(byt), "AppForm")
	}
}

func TestSearchResults_SearchByFieldValue(t *testing.T) {

	responseSwitch := 0
	testApiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch responseSwitch {
		case 0:
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":[{"id":33129,"notes":{},"fields":{},"title":"New AppForm Test"},{"id":33147,"notes":{},"fields":{},"title":"New AppForm Test"},{"id":33148,"notes":{},"fields":{},"title":"New AppForm Test"},{"id":33484,"notes":{},"fields":{},"title":"New AppForm Test"},{"id":33489,"notes":{},"fields":{},"title":"New AppForm Test"}],"additional_data":{"user":{"profile":{"id":0,"email":"","name":"","is_admin":false,"default_currency":"","icon_url":null,"activated":false},"locale":{"language":"","country":"","uses_12_hour_clock":false},"timezone":{"name":"","offset":0}},"multiple_companies":false,"default_company_id":0,"company_id":0,"since_timestamp":"","last_timestamp_on_page":"","pagination":{"start":0,"limit":100,"more_items_in_collection":false}}}`))
		default:
			t.Error("Too many requests")
		}
		responseSwitch = responseSwitch + 1
	}))

	config := &pipedrive.Config{
		APIKey:  "1",
		BaseURL: testApiServer.URL,
	}
	mockClient := pipedrive.NewClient(config)

	result, _, err := mockClient.SearchResults.SearchByFieldValue(context.Background(), &pipedrive.SearchResultsListOptions{
		Term:          "New AppForm Test",
		ExactMatch:    1,
		FieldType:     "dealField",
		FieldKey:      "title",
		ReturnItemIDs: 1,
	})

	if err != nil {
		t.Errorf("Could not get search results: %v", err)
	}

	if result.Success != true {
		t.Error("Got invalid search results")
	}

	for _, deal := range result.Data {
		assert.Equal(t, "New AppForm Test", deal.Title)
	}
}
