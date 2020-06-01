package pipedrive

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePerson(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(201)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":67,"company_id":7571105,"owner_id":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true,"value":11535881},"org_id":null,"name":"test test","first_name":"test","last_name":"test","open_deals_count":0,"related_open_deals_count":0,"closed_deals_count":0,"related_closed_deals_count":0,"participant_open_deals_count":0,"participant_closed_deals_count":0,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"reference_activities_count":0,"files_count":0,"notes_count":0,"followers_count":0,"won_deals_count":0,"related_won_deals_count":0,"lost_deals_count":0,"related_lost_deals_count":0,"active_flag":true,"phone":[{"label":"","value":"+61410036651","primary":true}],"email":[{"label":"","value":"test@test.com","primary":true}],"first_char":"a","update_time":"2020-05-31 23:01:23","add_time":"2020-05-31 23:01:23","visible_to":"3","picture_id":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"org_name":null,"cc_email":"testcompany151@pipedrivemail.com","owner_name":"Tom Shi"},"related_objects":{"user":{"11535881":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true}}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outPerson := &BasePersonObject{}
		err := testClient.CreatePerson(context.Background(), &BasePersonObject{}, &BaseResponse{Data: outPerson})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outPerson) {
			assert.NotEmpty(t, outPerson.Name)
		}
	})

	t.Run("Test handle 400 bad request", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Name must be given.","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outPerson := &BasePersonObject{}
		err := testClient.CreatePerson(context.Background(), &BasePersonObject{}, &BaseResponse{Data: outPerson})
		if assert.NotNil(t, err) {
			assert.Equal(t, "POST: 400 \"Name must be given.\"", err.Error())
		}
	})
}

func TestUpdatePerson(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":67,"company_id":7571105,"owner_id":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true,"value":11535881},"org_id":null,"name":"test test","first_name":"test","last_name":"test","open_deals_count":0,"related_open_deals_count":0,"closed_deals_count":0,"related_closed_deals_count":0,"participant_open_deals_count":0,"participant_closed_deals_count":0,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"reference_activities_count":0,"files_count":0,"notes_count":0,"followers_count":0,"won_deals_count":0,"related_won_deals_count":0,"lost_deals_count":0,"related_lost_deals_count":0,"active_flag":true,"phone":[{"label":"","value":"+61410036651","primary":true}],"email":[{"label":"","value":"test@test.com","primary":true}],"first_char":"a","update_time":"2020-05-31 23:01:23","add_time":"2020-05-31 23:01:23","visible_to":"3","picture_id":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"org_name":null,"cc_email":"testcompany151@pipedrivemail.com","owner_name":"Tom Shi"},"related_objects":{"user":{"11535881":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true}}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outPerson := &BasePersonObject{}
		err := testClient.UpdatePerson(context.Background(), 1, &BasePersonObject{}, &BaseResponse{Data: outPerson})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outPerson) {
			assert.NotEmpty(t, outPerson.Name)
		}
	})

	t.Run("Test handle 400 bad request", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Bad request","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outPerson := &BasePersonObject{}
		err := testClient.UpdatePerson(context.Background(), 1, &BasePersonObject{}, &BaseResponse{Data: outPerson})
		if assert.NotNil(t, err) {
			assert.Equal(t, "PUT: 400 \"Bad request\"", err.Error())
		}
	})
}

func TestDeletePersons(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":[69,70]}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		err := testClient.DeletePersons(context.Background(), []int{69, 70})
		assert.Nil(t, err)
	})
}

func TestDeletePerson(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":67}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		err := testClient.DeletePerson(context.Background(), 67)
		assert.Nil(t, err)
	})

	t.Run("Test handle 400 bad request", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Please provide at least one item to delete","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		err := testClient.DeletePerson(context.Background(), 1)
		if assert.NotNil(t, err) {
			assert.Equal(t, "DELETE: 400 \"Please provide at least one item to delete\"", err.Error())
		}
	})
}

func TestSearchPersons(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"items":[{"result_score":0.29955,"item":{"id":60,"type":"person","name":"test test","phones":["0410036651"],"emails":["test@test.com"],"visible_to":3,"owner":{"id":11535881},"organization":null,"custom_fields":[],"notes":[]}},{"result_score":0.29955,"item":{"id":9,"type":"person","name":"test test","phones":["040000000","0410036651"],"emails":[],"visible_to":3,"owner":{"id":11535881},"organization":null,"custom_fields":[],"notes":[]}},{"result_score":0.29955,"item":{"id":7,"type":"person","name":"test test","phones":["0410036651"],"emails":[],"visible_to":3,"owner":{"id":11535881},"organization":null,"custom_fields":[],"notes":[]}}]},"additional_data":{"pagination":{"start":0,"limit":100,"more_items_in_collection":false}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		searchOpt := &SearchPersonsOptions{
			Term: "0400000000",
		}
		out, err := testClient.SearchPersons(context.Background(), searchOpt)
		if err != nil {
			t.Fatal(err)
		}
		if assert.NotEmpty(t, out.Data.Items) {
			for _, item := range out.Data.Items {
				assert.NotNil(t, item.Person.Name)
			}
		}

	})

	t.Run("Test handle empty success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"items":[]},"additional_data":{"pagination":{"start":0,"limit":100,"more_items_in_collection":false}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		searchOpt := &SearchPersonsOptions{
			Term: "0400000000",
		}
		out, err := testClient.SearchPersons(context.Background(), searchOpt)
		if err != nil {
			t.Fatal(err)
		}

		assert.Len(t, out.Data.Items, 0)

	})

	t.Run("Test handle 400 bad request", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":{"message":"ServerError [ERR_INVALID_INPUT]: Invalid input: \"term\" is required","code":400,"failed_fields":["term"]}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		_, err := testClient.SearchPersons(context.Background(), &SearchPersonsOptions{
			Term: "0400000000",
		})
		if assert.NotNil(t, err) {
			assert.Equal(t, "GET: 400 {\"message\":\"ServerError [ERR_INVALID_INPUT]: Invalid input: \\\"term\\\" is required\",\"code\":400,\"failed_fields\":[\"term\"]}", err.Error())
		}
	})
}

func TestGetPerson(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":3,"company_id":7571105,"owner_id":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true,"value":11535881},"org_id":null,"name":"testtest","first_name":"test test","last_name":"test test","open_deals_count":4,"related_open_deals_count":0,"closed_deals_count":0,"related_closed_deals_count":0,"participant_open_deals_count":0,"participant_closed_deals_count":0,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"reference_activities_count":0,"files_count":0,"notes_count":0,"followers_count":1,"won_deals_count":0,"related_won_deals_count":0,"lost_deals_count":0,"related_lost_deals_count":0,"active_flag":true,"phone":[{"value":"","primary":true}],"email":[{"value":"","primary":true}],"first_char":"t","update_time":"2020-06-01 01:26:38","add_time":"2020-05-25 04:15:05","visible_to":"3","picture_id":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"f9c0ef7fa07cd385b75ab6b633554c306bc46a25":null,"org_name":null,"cc_email":"testcompany151@pipedrivemail.com","owner_name":"Tom Shi"},"additional_data":{"dropbox_email":"testcompany151@pipedrivemail.com"},"related_objects":{"user":{"11535881":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true}}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outPerson := &BasePersonObject{}
		err := testClient.GetPerson(context.Background(), 1, &BaseResponse{Data: outPerson})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outPerson) {
			assert.NotEmpty(t, outPerson.Name)
		}
	})

	t.Run("Test handle 404 not found", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Person not found","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outPerson := &BasePersonObject{}
		err := testClient.GetPerson(context.Background(), 1, &BaseResponse{Data: outPerson})
		if assert.NotNil(t, err) {
			assert.Equal(t, "GET: 404 \"Person not found\"", err.Error())
		}
	})
}
