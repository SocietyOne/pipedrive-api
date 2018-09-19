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

func TestDeal_Create(t *testing.T) {

	responseSwitch := 0
	testApiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch responseSwitch {
		case 0:
			w.WriteHeader(201)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":33489,"creator_user_id":{"id":3141486,"name":"Outbound Quotes","email":"verosupport@societyone.com.au","has_pic":true,"pic_hash":"5491de99d3bb382b8fa14b73377cea00","active_flag":true,"value":3141486},"user_id":{"id":3141486,"name":"Outbound Quotes","email":"verosupport@societyone.com.au","has_pic":true,"pic_hash":"5491de99d3bb382b8fa14b73377cea00","active_flag":true,"value":3141486},"person_id":{"name":"Tom New APITest","email":[{"label":"","value":"NewAPITest@Mail.com","primary":true}],"phone":[{"label":"","value":"0400000000","primary":true}],"value":27751},"org_id":null,"stage_id":15,"title":"New AppForm Test","value":543210,"currency":"AUD","add_time":"2018-09-17 06:13:10","update_time":"2018-09-17 06:13:10","stage_change_time":null,"active":true,"deleted":false,"status":"open","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":null,"pipeline_id":5,"won_time":null,"first_won_time":null,"lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":0,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"reference_activities_count":0,"participants_count":0,"expected_close_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"stage_order_nr":1,"person_name":"Tom New APITest","org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"$543,210","weighted_value":543210,"formatted_weighted_value":"$543,210","weighted_value_currency":"AUD","rotten_time":null,"owner_name":"Outbound Quotes","cc_email":"societyoneoutbound+deal33489@pipedrivemail.com","org_hidden":false,"person_hidden":false},"additional_data":{"user":{"profile":{"id":0,"email":"","name":"","is_admin":false,"default_currency":"","icon_url":null,"activated":false},"locale":{"language":"","country":"","uses_12_hour_clock":false},"timezone":{"name":"","offset":0}},"multiple_companies":false,"default_company_id":0,"company_id":0,"since_timestamp":"","last_timestamp_on_page":"","pagination":{"start":0,"limit":0,"more_items_in_collection":false}},"related_objects":{"person":{"27751":{"email":[{"label":"","primary":true,"value":"NewAPITest@Mail.com"}],"id":27751,"name":"Tom New APITest","phone":[{"label":"","primary":true,"value":"0400000000"}]}},"user":{"3141486":{"active_flag":true,"email":"verosupport@societyone.com.au","has_pic":true,"id":3141486,"name":"Outbound Quotes","pic_hash":"5491de99d3bb382b8fa14b73377cea00"}}}}`))
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

	deal := &pipedrive.Deal{

		Title:    "New AppForm Test",
		Value:    543210,
		StageID:  15,
		PersonID: 27751,
		UserID:   3141486, //Outbound Quotes

	}

	response, httpResponse, err := mockClient.Deals.Create(context.Background(), deal)

	if err != nil {
		t.Errorf("Invalid create deal request: %v", err)
	}

	fmt.Println(httpResponse.StatusCode)
	fmt.Println(string(response.Data))

	var createdDeal pipedrive.Deal
	err = json.NewDecoder(bytes.NewReader(response.Data)).Decode(&createdDeal)

	if response.Success != true {
		t.Error("Could not create deal")
	}

	assert.Equal(t, "New AppForm Test", createdDeal.Title, "Deal title did not match expected")
}

func TestDeal_Update(t *testing.T) {

	responseSwitch := 0
	testApiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch responseSwitch {
		case 0:
			w.WriteHeader(201)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":31706,"creator_user_id":{"id":3141486,"name":"Outbound Quotes","email":"verosupport@societyone.com.au","has_pic":true,"pic_hash":"5491de99d3bb382b8fa14b73377cea00","active_flag":true,"value":3141486},"user_id":{"id":3141486,"name":"Outbound Quotes","email":"verosupport@societyone.com.au","has_pic":true,"pic_hash":"5491de99d3bb382b8fa14b73377cea00","active_flag":true,"value":3141486},"person_id":{"name":"Tom New APITest","email":[{"label":"","value":"NewAPITest@Mail.com","primary":true}],"phone":[{"label":"","value":"0400000000","primary":true}],"value":27751},"org_id":null,"stage_id":16,"title":"New APP Form Test Deal","value":666,"currency":"AUD","add_time":"2018-08-30 06:03:31","update_time":"2018-09-17 06:08:11","stage_change_time":"2018-09-17 06:08:11","active":true,"deleted":false,"status":"open","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":null,"pipeline_id":5,"won_time":null,"first_won_time":null,"lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":1,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"reference_activities_count":0,"participants_count":1,"expected_close_date":null,"last_incoming_mail_time":null,"stage_order_nr":2,"person_name":"Tom New APITest","org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"$666","weighted_value":666,"formatted_weighted_value":"$666","weighted_value_currency":"AUD","rotten_time":null,"owner_name":"Outbound Quotes","cc_email":"societyoneoutbound+deal31706@pipedrivemail.com","org_hidden":false,"person_hidden":false},"additional_data":{"user":{"profile":{"id":0,"email":"","name":"","is_admin":false,"default_currency":"","icon_url":null,"activated":false},"locale":{"language":"","country":"","uses_12_hour_clock":false},"timezone":{"name":"","offset":0}},"multiple_companies":false,"default_company_id":0,"company_id":0,"since_timestamp":"","last_timestamp_on_page":"","pagination":{"start":0,"limit":0,"more_items_in_collection":false}},"related_objects":{"person":{"27751":{"email":[{"label":"","primary":true,"value":"NewAPITest@Mail.com"}],"id":27751,"name":"Tom New APITest","phone":[{"label":"","primary":true,"value":"0400000000"}]}},"user":{"3141486":{"active_flag":true,"email":"verosupport@societyone.com.au","has_pic":true,"id":3141486,"name":"Outbound Quotes","pic_hash":"5491de99d3bb382b8fa14b73377cea00"}}}}`))
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

	deal := &pipedrive.Deal{

		Value:   666,
		StageID: 16,
	}
	dealIDToUpdate := 31706

	response, httpResponse, err := mockClient.Deals.Update(context.Background(), dealIDToUpdate, deal)

	if err != nil {
		t.Errorf("Invalid update deal request: %v", err)
	}

	fmt.Println(httpResponse.StatusCode)
	fmt.Println(string(response.Data))

	var updatedDeal pipedrive.Deal
	err = json.NewDecoder(bytes.NewReader(response.Data)).Decode(&updatedDeal)

	if response.Success != true {
		t.Error("Could not update deal")
	}

	assert.Equal(t, dealIDToUpdate, updatedDeal.ID, "Get Deal ID and Expected are not equal")
}

func TestDeal_Get(t *testing.T) {
	responseSwitch := 0
	testApiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		switch responseSwitch {
		case 0:
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":31706,"creator_user_id":{"id":3141486,"name":"Outbound Quotes","email":"verosupport@societyone.com.au","has_pic":true,"pic_hash":"5491de99d3bb382b8fa14b73377cea00","active_flag":true,"value":3141486},"user_id":{"id":3141486,"name":"Outbound Quotes","email":"verosupport@societyone.com.au","has_pic":true,"pic_hash":"5491de99d3bb382b8fa14b73377cea00","active_flag":true,"value":3141486},"person_id":{"name":"Tom New APITest","email":[{"label":"","value":"NewAPITest@Mail.com","primary":true}],"phone":[{"label":"","value":"0400000000","primary":true}],"value":27751},"org_id":null,"stage_id":16,"title":"New APP Form Test Deal","value":666,"currency":"AUD","add_time":"2018-08-30 06:03:31","update_time":"2018-09-14 01:56:46","stage_change_time":"2018-09-14 01:56:46","active":true,"deleted":false,"status":"open","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":null,"pipeline_id":5,"won_time":null,"first_won_time":null,"lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":1,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"reference_activities_count":0,"participants_count":1,"expected_close_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"stage_order_nr":2,"person_name":"Tom New APITest","org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"$666","weighted_value":666,"formatted_weighted_value":"$666","weighted_value_currency":"AUD","rotten_time":null,"owner_name":"Outbound Quotes","cc_email":"societyoneoutbound+deal31706@pipedrivemail.com","org_hidden":false,"person_hidden":false,"average_time_to_won":{"y":0,"m":0,"d":0,"h":0,"i":0,"s":0,"total_seconds":0},"average_stage_progress":0,"age":{"y":0,"m":0,"d":18,"h":0,"i":1,"s":22,"total_seconds":1555282},"stay_in_pipeline_stages":{"times_in_stages":{"15":1281138,"16":274139,"17":5},"order_of_stages":[15,16,17]},"last_activity":null,"next_activity":null},"additional_data":{"user":{"profile":{"id":0,"email":"","name":"","is_admin":false,"default_currency":"","icon_url":null,"activated":false},"locale":{"language":"","country":"","uses_12_hour_clock":false},"timezone":{"name":"","offset":0}},"multiple_companies":false,"default_company_id":0,"company_id":0,"since_timestamp":"","last_timestamp_on_page":"","pagination":{"start":0,"limit":0,"more_items_in_collection":false}},"related_objects":{"person":{"27751":{"email":[{"label":"","primary":true,"value":"NewAPITest@Mail.com"}],"id":27751,"name":"Tom New APITest","phone":[{"label":"","primary":true,"value":"0400000000"}]}},"user":{"3141486":{"active_flag":true,"email":"verosupport@societyone.com.au","has_pic":true,"id":3141486,"name":"Outbound Quotes","pic_hash":"5491de99d3bb382b8fa14b73377cea00"}}}}`))
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

	id := 31706

	response, httpResponse, err := mockClient.Deals.Get(context.Background(), id)

	if err != nil {
		t.Errorf("Invalid get deal request: %v", err)
	}

	fmt.Println(httpResponse.StatusCode)
	fmt.Println(string(response.Data))

	var deal pipedrive.Deal
	err = json.NewDecoder(bytes.NewReader(response.Data)).Decode(&deal)

	if err != nil {
		t.Error("Failed to decode")
	}

	if response.Success != true {
		t.Error("Could not get deal")
	}

	assert.Equal(t, id, deal.ID, "Get Deal ID and Expected are not equal")

}
