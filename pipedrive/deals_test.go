package pipedrive

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDeal(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(201)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":5,"creator_user_id":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true,"value":11535881},"user_id":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true,"value":11535881},"person_id":{"active_flag":true,"name":"testtest","email":[{"value":"","primary":true}],"phone":[{"value":"","primary":true}],"value":3},"org_id":null,"stage_id":1,"title":"test 456","value":0,"currency":"AUD","add_time":"2020-06-01 02:41:35","update_time":"2020-06-01 02:41:35","stage_change_time":null,"active":true,"deleted":false,"status":"open","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":null,"pipeline_id":1,"won_time":null,"first_won_time":null,"lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":0,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"reference_activities_count":0,"participants_count":0,"expected_close_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"f68bc64c61ed5be74939265930336b9424d7c39b":"custom","stage_order_nr":0,"person_name":"testtest","org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"A$0","weighted_value":0,"formatted_weighted_value":"A$0","weighted_value_currency":"AUD","rotten_time":null,"owner_name":"Tom Shi","cc_email":"testcompany151+deal5@pipedrivemail.com","org_hidden":false,"person_hidden":false},"related_objects":{"user":{"11535881":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true}},"person":{"3":{"active_flag":true,"id":3,"name":"testtest","email":[{"value":"","primary":true}],"phone":[{"value":"","primary":true}]}}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outDeal := &BaseDealObject{}
		err := testClient.CreateDeal(context.Background(), &BaseDealObject{}, &BaseResponse{Data: outDeal})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outDeal) {
			assert.NotEmpty(t, outDeal.Title)
		}
	})

	t.Run("Test handle 400 bad request", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Invalid field(s) in the payload: asdasdasd","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outDeal := &BaseDealObject{}
		err := testClient.CreateDeal(context.Background(), &BaseDealObject{}, &BaseResponse{Data: outDeal})
		if assert.NotNil(t, err) {
			assert.Equal(t, "POST: 400 \"Invalid field(s) in the payload: asdasdasd\"", err.Error())
		}
	})
}

func TestUpdateDeal(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":1,"creator_user_id":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true,"value":11535881},"user_id":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true,"value":11535881},"person_id":{"active_flag":true,"name":"testtest","email":[{"value":"","primary":true}],"phone":[{"value":"","primary":true}],"value":3},"org_id":null,"stage_id":1,"title":"123","value":0,"currency":"AUD","add_time":"2020-06-01 00:59:46","update_time":"2020-06-01 02:49:25","stage_change_time":null,"active":false,"deleted":false,"status":"open","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":"2020-06-01 02:36:31","pipeline_id":1,"won_time":null,"first_won_time":null,"lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":1,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"reference_activities_count":0,"participants_count":1,"expected_close_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"f68bc64c61ed5be74939265930336b9424d7c39b":null,"stage_order_nr":0,"person_name":"testtest","org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"A$0","weighted_value":0,"formatted_weighted_value":"A$0","weighted_value_currency":"AUD","rotten_time":null,"owner_name":"Tom Shi","cc_email":"testcompany151+deal1@pipedrivemail.com","org_hidden":false,"person_hidden":false},"related_objects":{"user":{"11535881":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true}},"person":{"3":{"active_flag":true,"id":3,"name":"testtest","email":[{"value":"","primary":true}],"phone":[{"value":"","primary":true}]}}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outDeal := &BaseDealObject{}
		err := testClient.UpdateDeal(context.Background(), 1, &BaseDealObject{}, &BaseResponse{Data: outDeal})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outDeal) {
			assert.NotEmpty(t, outDeal.Title)
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

		outDeal := &BaseDealObject{}
		err := testClient.UpdateDeal(context.Background(), 1, &BaseDealObject{}, &BaseResponse{Data: outDeal})
		if assert.NotNil(t, err) {
			assert.Equal(t, "PUT: 400 \"Bad request\"", err.Error())
		}
	})
}

func TestDeleteDeals(t *testing.T) {
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

		err := testClient.DeleteDeals(context.Background(), []int{69, 70})
		assert.Nil(t, err)
	})
}

func TestDeleteDeal(t *testing.T) {
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

		err := testClient.DeleteDeal(context.Background(), 67)
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

		err := testClient.DeleteDeal(context.Background(), 1)
		if assert.NotNil(t, err) {
			assert.Equal(t, "DELETE: 400 \"Please provide at least one item to delete\"", err.Error())
		}
	})
}

func TestSearchDeals(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"items":[{"result_score":1.1688,"item":{"id":5,"type":"deal","title":"test 456","value":null,"currency":"AUD","status":"open","visible_to":3,"owner":{"id":11535881},"stage":{"id":1,"name":"Lead In"},"person":{"id":3,"name":"testtest"},"organization":null,"custom_fields":["custom"],"notes":[]}},{"result_score":1.1688,"item":{"id":4,"type":"deal","title":"test 456","value":null,"currency":"AUD","status":"open","visible_to":3,"owner":{"id":11535881},"stage":{"id":1,"name":"Lead In"},"person":{"id":3,"name":"testtest"},"organization":null,"custom_fields":["custom"],"notes":[]}}]},"additional_data":{"pagination":{"start":0,"limit":100,"more_items_in_collection":false}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		searchOpt := &SearchDealsOptions{
			Term: "0400000000",
		}
		out, err := testClient.SearchDeals(context.Background(), searchOpt)
		if err != nil {
			t.Fatal(err)
		}
		if assert.NotEmpty(t, out.Data.Items) {
			for _, item := range out.Data.Items {
				assert.NotNil(t, item.Deal.Title)
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

		searchOpt := &SearchDealsOptions{
			Term: "0400000000",
		}
		out, err := testClient.SearchDeals(context.Background(), searchOpt)
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
			w.Write([]byte(`{"success":false,"error":{"message":"ServerError [ERR_INVALID_INPUT]: Invalid input: \"term\" is not allowed to be empty","code":400,"failed_fields":["term"]}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		_, err := testClient.SearchDeals(context.Background(), &SearchDealsOptions{
			Term: "0400000000",
		})
		if assert.NotNil(t, err) {
			assert.Equal(t, "GET: 400 {\"message\":\"ServerError [ERR_INVALID_INPUT]: Invalid input: \\\"term\\\" is not allowed to be empty\",\"code\":400,\"failed_fields\":[\"term\"]}", err.Error())
		}
	})
}

func TestGetDeal(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":1,"creator_user_id":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true,"value":11535881},"user_id":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true,"value":11535881},"person_id":{"active_flag":true,"name":"testtest","email":[{"value":"","primary":true}],"phone":[{"value":"","primary":true}],"value":3},"org_id":null,"stage_id":1,"title":"123","value":0,"currency":"AUD","add_time":"2020-06-01 00:59:46","update_time":"2020-06-01 02:49:25","stage_change_time":null,"active":false,"deleted":false,"status":"open","probability":null,"next_activity_date":null,"next_activity_time":null,"next_activity_id":null,"last_activity_id":null,"last_activity_date":null,"lost_reason":null,"visible_to":"3","close_time":"2020-06-01 02:36:31","pipeline_id":1,"won_time":null,"first_won_time":null,"lost_time":null,"products_count":0,"files_count":0,"notes_count":0,"followers_count":1,"email_messages_count":0,"activities_count":0,"done_activities_count":0,"undone_activities_count":0,"reference_activities_count":0,"participants_count":1,"expected_close_date":null,"last_incoming_mail_time":null,"last_outgoing_mail_time":null,"label":null,"f68bc64c61ed5be74939265930336b9424d7c39b":null,"stage_order_nr":0,"person_name":"testtest","org_name":null,"next_activity_subject":null,"next_activity_type":null,"next_activity_duration":null,"next_activity_note":null,"formatted_value":"A$0","weighted_value":0,"formatted_weighted_value":"A$0","weighted_value_currency":"AUD","rotten_time":null,"owner_name":"Tom Shi","cc_email":"testcompany151+deal1@pipedrivemail.com","org_hidden":false,"person_hidden":false,"average_time_to_won":{"y":0,"m":0,"d":0,"h":0,"i":0,"s":0,"total_seconds":0},"average_stage_progress":0,"age":{"y":0,"m":0,"d":0,"h":1,"i":53,"s":21,"total_seconds":6801},"stay_in_pipeline_stages":{"times_in_stages":{"1":6801},"order_of_stages":[1]},"last_activity":null,"next_activity":null},"additional_data":{"dropbox_email":"testcompany151+deal1@pipedrivemail.com"},"related_objects":{"user":{"11535881":{"id":11535881,"name":"Tom Shi","email":"tom.shi@societyone.com.au","has_pic":0,"pic_hash":null,"active_flag":true}},"person":{"3":{"active_flag":true,"id":3,"name":"testtest","email":[{"value":"","primary":true}],"phone":[{"value":"","primary":true}]}}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outDeal := &BaseDealObject{}
		err := testClient.GetDeal(context.Background(), 1, &BaseResponse{Data: outDeal})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outDeal) {
			assert.NotEmpty(t, outDeal.Title)
		}
	})

	t.Run("Test handle 404 not found", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Deal not found","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outDeal := &BaseDealObject{}
		err := testClient.GetDeal(context.Background(), 1, &BaseResponse{Data: outDeal})
		if assert.NotNil(t, err) {
			assert.Equal(t, "GET: 404 \"Deal not found\"", err.Error())
		}
	})
}
