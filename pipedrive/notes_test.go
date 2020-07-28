package pipedrive

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(201)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":1,"user_id":11535881,"deal_id":1,"person_id":3,"org_id":null,"content":"test by tom","add_time":"2020-06-01 03:27:05","update_time":"2020-06-01 03:27:05","active_flag":true,"pinned_to_deal_flag":false,"pinned_to_person_flag":false,"pinned_to_organization_flag":false,"last_update_user_id":null,"organization":null,"person":{"name":"testtest"},"deal":{"title":"123"},"user":{"email":"tom.shi@societyone.com.au","name":"Tom Shi","icon_url":null,"is_you":true}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outNote := &BaseNoteObject{}
		err := testClient.CreateNote(context.Background(), &BaseNoteObject{}, &BaseResponse{Data: outNote})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outNote) {
			assert.NotEmpty(t, outNote.Content)
		}
	})

	t.Run("Test handle 400 bad request", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Cannot add/update note because deal does not exist or you have no permissions to access this deal.","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outNote := &BaseNoteObject{}
		err := testClient.CreateNote(context.Background(), &BaseNoteObject{}, &BaseResponse{Data: outNote})
		if assert.NotNil(t, err) {
			assert.Equal(t, "POST: 400 \"Cannot add/update note because deal does not exist or you have no permissions to access this deal.\"", err.Error())
		}
	})
}

func TestUpdateNote(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":1,"user_id":11535881,"deal_id":1,"person_id":3,"org_id":null,"content":"asd","add_time":"2020-06-01 03:27:05","update_time":"2020-06-01 04:33:41","active_flag":true,"pinned_to_deal_flag":false,"pinned_to_person_flag":false,"pinned_to_organization_flag":false,"last_update_user_id":11535881,"organization":null,"person":{"name":"testtest"},"deal":{"title":"123"},"user":{"email":"tom.shi@societyone.com.au","name":"Tom Shi","icon_url":null,"is_you":true}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outNote := &BaseNoteObject{}
		err := testClient.UpdateNote(context.Background(), 1, &BaseNoteObject{}, &BaseResponse{Data: outNote})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outNote) {
			assert.NotEmpty(t, outNote.Content)
		}
	})

	t.Run("Test handle 400 bad request", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(400)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Note needs to have a content.","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outNote := &BaseNoteObject{}
		err := testClient.UpdateNote(context.Background(), 1, &BaseNoteObject{}, &BaseResponse{Data: outNote})
		if assert.NotNil(t, err) {
			assert.Equal(t, "PUT: 400 \"Note needs to have a content.\"", err.Error())
		}
	})
}

func TestDeleteNote(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":true}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		err := testClient.DeleteNote(context.Background(), 1)
		assert.Nil(t, err)
	})

	t.Run("Test handle 410 gone", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(410)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":true}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		err := testClient.DeleteNote(context.Background(), 1)
		assert.Nil(t, err)
	})
}

func TestGetNote(t *testing.T) {
	t.Run("Test handle success response", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":true,"data":{"id":1,"company_id":7571105,"user_id":11535881,"deal_id":1,"person_id":3,"org_id":null,"content":"abc","add_time":"2020-06-01 03:27:05","update_time":"2020-06-01 04:36:47","active_flag":false,"pinned_to_deal_flag":false,"pinned_to_person_flag":false,"pinned_to_organization_flag":false,"last_update_user_id":11535881,"organization":null,"person":{"name":"testtest"},"deal":{"title":"123"},"user":{"email":"tom.shi@societyone.com.au","name":"Tom Shi","icon_url":null,"is_you":true}}}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outNote := &BaseNoteObject{}
		err := testClient.GetNote(context.Background(), 1, &BaseResponse{Data: outNote})
		if err != nil {
			t.Fatal(err)
		}

		if assert.NotNil(t, outNote) {
			assert.NotEmpty(t, outNote.Content)
		}
	})

	t.Run("Test handle 404 not found", func(t *testing.T) {
		testAPI := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// Canned Response
			w.WriteHeader(404)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"success":false,"error":"Note not found","error_info":"Please check developers.pipedrive.com for more information about Pipedrive API.","data":null,"additional_data":null}`))
		}))
		config := &Config{
			APIKey:  "1",
			BaseURL: testAPI.URL,
		}
		testClient := NewClient(config)

		outNote := &BaseNoteObject{}
		err := testClient.GetNote(context.Background(), 1, &BaseResponse{Data: outNote})
		if assert.NotNil(t, err) {
			assert.Equal(t, "GET: 404 \"Note not found\"", err.Error())
		}
	})
}
