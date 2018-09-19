package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SocietyOne/pipedrive-api/pipedrive"
	"github.com/stretchr/testify/assert"
)

func TestPersons_Find(t *testing.T) {
	responseSwitch := 0
	testApiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch responseSwitch {
		case 0:
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":[{"id":27751,"name":"Tom New APITest","email":"NewAPITest@Mail.com","phone":"0400000000","org_id":null,"org_name":"","visible_to":"3","picture":null}],"additional_data":{"user":{"profile":{"id":0,"email":"","name":"","is_admin":false,"default_currency":"","icon_url":null,"activated":false},"locale":{"language":"","country":"","uses_12_hour_clock":false},"timezone":{"name":"","offset":0}},"multiple_companies":false,"default_company_id":0,"company_id":0,"since_timestamp":"","last_timestamp_on_page":"","pagination":{"start":0,"limit":100,"more_items_in_collection":false}}}`))
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

	searchTerm := "Tom New APITest"
	response, httpResponse, err := mockClient.Persons.Find(context.Background(), searchTerm)

	if err != nil {
		t.Errorf("Could not get search results: %v", err)
	}

	fmt.Println(httpResponse.StatusCode)
	fmt.Println(string(response.Data))

	var persons []pipedrive.Person
	err = json.NewDecoder(bytes.NewReader(response.Data)).Decode(&persons)

	if response.Success != true {
		t.Error("Got invalid search results")
	}

	assert.Equal(t, searchTerm, persons[0].Name, "Response name does not match search term")
}

func TestPersons_FindByEmail(t *testing.T) {
	responseSwitch := 0
	testApiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch responseSwitch {
		case 0:
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":[{"id":27751,"name":"Tom New APITest","email":"NewAPITest@Mail.com","phone":"0400000000","org_id":null,"org_name":"","visible_to":"3","additional_emails":[]}],"additional_data":{"user":{"profile":{"id":0,"email":"","name":"","is_admin":false,"default_currency":"","icon_url":null,"activated":false},"locale":{"language":"","country":"","uses_12_hour_clock":false},"timezone":{"name":"","offset":0}},"multiple_companies":false,"default_company_id":0,"company_id":0,"since_timestamp":"","last_timestamp_on_page":"","pagination":{"start":0,"limit":100,"more_items_in_collection":false}}}`))
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

	email := "NewAPITest@Mail.com"
	response, httpResponse, err := mockClient.Persons.FindByEmail(context.Background(), email)

	if err != nil {
		t.Errorf("Could not get search results: %v", err)
	}

	fmt.Println(httpResponse.StatusCode)
	fmt.Println(string(response.Data))

	var persons []pipedrive.Person
	err = json.NewDecoder(bytes.NewReader(response.Data)).Decode(&persons)

	if response.Success != true {
		t.Error("Got invalid search results")
	}

	assert.Equal(t, email, persons[0].Email, "Response email does not match search email")
}
