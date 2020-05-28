package test

import (
	"context"
	"fmt"
	"os"
	"reflect"
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

func TestPersons(t *testing.T) {

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

func PrintType(in pipedrive.Person) {

	interfacePerson := in.(interface{})

	reflectType := reflect.TypeOf(interfacePerson)
	reflectKind := reflectType.Kind()
	reflectValue := reflect.ValueOf(interfacePerson)

	fmt.Println(reflectType)
	fmt.Println(reflectKind)
	fmt.Println(reflectValue)
}
