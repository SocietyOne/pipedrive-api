package test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/SocietyOne/pipedrive-api/pipedrive"
)

type TestDealObject struct {
	pipedrive.BaseDealObject
	Value *float64 `json:"value,omitempty"`
}

type TestPerson struct {
	pipedrive.BasePersonObject
	FirstChar *string `json:"first_char,omitempty"`
}

func TestIntegration(t *testing.T) {

	os.Setenv("PIPEDRIVE_API_TOKEN", "487322921050ae1268e610da9c6be5a1759a16de") // This is a trial company token

	apiKey := os.Getenv("PIPEDRIVE_API_TOKEN")
	client := pipedrive.NewClient(pipedrive.NewConfig(apiKey))

	personName := "test1"
	randomChar := "a"
	personToCreate := &TestPerson{
		BasePersonObject: pipedrive.BasePersonObject{
			Name: &personName,
			OrgID: &pipedrive.OrgID{
				ID: 1,
			},
		},
		FirstChar: &randomChar,
	}
	createdPerson := &TestPerson{}
	response := &pipedrive.BaseResponse{
		Data: createdPerson,
	}
	err := client.CreatePerson(context.Background(), personToCreate, response)
	if err != nil {
		t.Fatal(err)
	}
	person := *createdPerson
	fmt.Println(person)
}
