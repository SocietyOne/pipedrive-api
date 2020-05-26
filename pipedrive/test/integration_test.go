package test

import (
	"context"
	"encoding/json"
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
	FName *string `json:"first_name,omitempty"`
}

//MarshalJSON is a Marshalling override
// func (o TestPerson) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(o)
// }

//UnmarshalJSON is an unmarshalling override
func (o *TestPerson) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, o)
}

func TestPersons(t *testing.T) {

	apiKey := os.Getenv("PIPEDRIVE_API_TOKEN")
	client := pipedrive.NewClient(pipedrive.NewConfig(apiKey))

	personName := "test1"
	randomFirstName := "abcabc"
	personToCreate := TestPerson{
		BasePersonObject: pipedrive.BasePersonObject{
			Name: &personName,
			// OrgID: &pipedrive.OrgID{
			// 	ID: 1,
			// },
		},
		FName: &randomFirstName,
	}

	// PrintType(personToCreate)

	// personToCreate := pipedrive.BasePersonObject{
	// 	Name: &personName,
	// 	// OrgID: &pipedrive.OrgID{
	// 	// 	ID: 1,
	// 	// },
	// }

	createdPerson, err := client.CreatePerson(context.Background(), personToCreate)
	if err != nil {
		t.Fatal(err)
	}

	person := *createdPerson
	// testPerson := person.(TestPerson)

	fmt.Println(person)

	// getPerson, err := client.GetPerson(context.Background(), createdPerson.ID)
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
