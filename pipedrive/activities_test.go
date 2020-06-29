package pipedrive

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateActivity(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(201)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":57735,"company_id":1792874,"user_id":11535881,"done":false,"type":"call","reference_type":null,"reference_id":null,"conference_meeting_client":null,"conference_meeting_url":null,"due_date":"2020-06-29","due_time":"","duration":"","busy_flag":null,"add_time":"2020-06-29 04:31:07","marked_as_done_time":"","last_notification_time":null,"last_notification_user_id":null,"notification_language_id":null,"subject":"test subject","public_description":null,"calendar_sync_include_context":null,"location":null,"org_id":null,"person_id":11140,"deal_id":16143,"lead_id":null,"lead_title":"","active_flag":true,"update_time":"2020-06-29 04:31:07","update_user_id":null,"gcal_event_id":null,"google_calendar_id":null,"google_calendar_etag":null,"source_timezone":null,"rec_rule":null,"rec_rule_extension":null,"rec_master_activity_id":null,"conference_meeting_id":null,"note":null,"created_by_user_id":11535881,"location_subpremise":null,"location_street_number":null,"location_route":null,"location_sublocality":null,"location_locality":null,"location_admin_area_level_1":null,"location_admin_area_level_2":null,"location_country":null,"location_postal_code":null,"location_formatted_address":null,"attendees":null,"participants":[{"person_id":11140,"primary_flag":true}],"series":null,"org_name":null,"person_name":"Mario Pereira","deal_title":"PRE-PROD-TEST [Multiple] 423428.02 John Doe Webform","owner_name":"Tom Shi","person_dropbox_bcc":"societyonecollections@pipedrivemail.com","deal_dropbox_bcc":"societyonecollections+deal16143@pipedrivemail.com","assigned_to_user_id":11535881,"file":null},"additional_data":{"updates_story_id":438044},"related_objects":{"person":{"11140":{"active_flag":true,"id":11140,"name":"Mario Pereira","email":[{"label":"","value":"201575mfp@gmail.com","primary":true}],"phone":[{"label":"","value":"0474351676","primary":true}]}},"deal":{"16143":{"id":16143,"title":"PRE-PROD-TEST [Multiple] 423428.02 John Doe Webform","status":"deleted","value":0,"currency":"AUD","stage_id":80,"pipeline_id":10}},"user":{"11535881":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true}}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outActivity := &BaseActivityObject{}
		err := testClient.CreateActivity(context.Background(), &BaseActivityObject{}, &BaseResponse{Data: outActivity})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outActivity) {
			assert.NotEmpty(t, outActivity.Subject)
		}
	})

	t.Run("Test handle 400 bad request", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Inserting ID with POST request is not allowed. If you have any questions, please contact Pipedrive support","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outActivity := &BaseActivityObject{}
		err := testClient.CreateActivity(context.Background(), &BaseActivityObject{}, &BaseResponse{Data: outActivity})
		if assert.NotNil(t, err) {
			assert.Equal(t, "POST: 400 \"Inserting ID with POST request is not allowed. If you have any questions, please contact Pipedrive support\"", err.Error())
		}
	})
}
