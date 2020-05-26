package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// PersonsService handles activities related
// methods of the Pipedrive API.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons
type PersonsService service

type Email struct {
	Label   string `json:"label"`
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

type Phone struct {
	Label   string `json:"label"`
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

// Person represents a Pipedrive person.
type Person interface {
	// APIName() string
	// MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

type BasePersonObject struct {
	// Unsettable Fields
	ID int `json:"id,omitempty" force:"id,omitempty"`

	// Settable Fields
	Name      *string `json:"name,omitempty" force:"name,omitempty"`             // Required
	FirstName *string `json:"first_name,omitempty" force:"first_name,omitempty"` // Optional
	LastName  *string `json:"last_name,omitempty" force:"last_name,omitempty"`   // Optional
	// Phone     []*Phone `json:"phone,omitempty" force:"phone,omitempty"`
	// Email     []*Email `json:"email,omitempty" force:"email,omitempty"`
	// OrgID     *OrgID   `json:"org_id,omitempty"`

	// Unused Fields
	// OwnerID                         interface{} `json:"owner_id,omitempty"`
	// CompanyID                       int         `json:"company_id,omitempty"`
	// OpenDealsCount                  int         `json:"open_deals_count,omitempty"`
	// RelatedOpenDealsCount           int         `json:"related_open_deals_count,omitempty"`
	// ClosedDealsCount                int         `json:"closed_deals_count,omitempty"`
	// RelatedClosedDealsCount         int         `json:"related_closed_deals_count,omitempty"`
	// ParticipantOpenDealsCount       int         `json:"participant_open_deals_count,omitempty"`
	// ParticipantClosedDealsCount     int         `json:"participant_closed_deals_count,omitempty"`
	// EmailMessagesCount              int         `json:"email_messages_count,omitempty"`
	// ActivitiesCount                 int         `json:"activities_count,omitempty"`
	// DoneActivitiesCount             int         `json:"done_activities_count,omitempty"`
	// UndoneActivitiesCount           int         `json:"undone_activities_count,omitempty"`
	// ReferenceActivitiesCount        int         `json:"reference_activities_count,omitempty"`
	// FilesCount                      int         `json:"files_count,omitempty"`
	// NotesCount                      int         `json:"notes_count,omitempty"`
	// FollowersCount                  int         `json:"followers_count,omitempty"`
	// WonDealsCount                   int         `json:"won_deals_count,omitempty"`
	// RelatedWonDealsCount            int         `json:"related_won_deals_count,omitempty"`
	// LostDealsCount                  int         `json:"lost_deals_count,omitempty"`
	// RelatedLostDealsCount           int         `json:"related_lost_deals_count,omitempty"`
	// ActiveFlag                      bool        `json:"active_flag,omitempty"`
	// FirstChar                       string      `json:"first_char,omitempty"`
	// UpdateTime                      string      `json:"update_time,omitempty"`
	// AddTime                         string      `json:"add_time,omitempty"`
	// VisibleTo                       string      `json:"visible_to,omitempty"`
	// PictureID                       interface{} `json:"picture_id,omitempty"`
	// NextActivityDate                interface{} `json:"next_activity_date,omitempty"`
	// NextActivityTime                interface{} `json:"next_activity_time,omitempty"`
	// NextActivityID                  interface{} `json:"next_activity_id,omitempty"`
	// LastActivityID                  int         `json:"last_activity_id,omitempty"`
	// LastActivityDate                string      `json:"last_activity_date,omitempty"`
	// TimelineLastActivityTime        interface{} `json:"timeline_last_activity_time,omitempty"`
	// TimelineLastActivityTimeByOwner interface{} `json:"timeline_last_activity_time_by_owner,omitempty"`
	// LastIncomingMailTime            interface{} `json:"last_incoming_mail_time,omitempty"`
	// LastOutgoingMailTime            interface{} `json:"last_outgoing_mail_time,omitempty"`
	// OrgName                         interface{} `json:"org_name,omitempty"`
	// OwnerName                       string      `json:"owner_name,omitempty"`
	// CcEmail                         string      `json:"cc_email,omitempty"`
}

func (p BasePersonObject) APIName() string {
	return "deals"
}

func (p BasePersonObject) String() string {
	return Stringify(p)
}

// PersonsResponse represents multiple persons response.
type PersonsResponse struct {
	Success        bool            `json:"success"`
	Data           []*Person       `json:"data"`
	AdditionalData *AdditionalData `json:"additional_data,omitempty"`
}

// PersonResponse represents single person response.
type PersonResponse struct {
	Success        bool            `json:"success"`
	Data           *Person         `json:"data"`
	AdditionalData *AdditionalData `json:"additional_data,omitempty"`
	// related objects
}

// CreatePerson create a new person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/post_persons
func (c *Client) CreatePerson(ctx context.Context, person Person) (*Person, error) {

	req, err := c.NewRequest(http.MethodPost, "/persons", nil, person)
	// req, err := c.NewRequest(http.MethodPost, "/persons", nil, person.(interface{}))

	if err != nil {
		return nil, err
	}

	var personResponse *PersonResponse

	_, err = c.Do(ctx, req, &personResponse)

	if err != nil {
		return nil, err
	}

	return personResponse.Data, nil
	// return person, nil
}

// UpdatePerson update a specific person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id
func (c *Client) UpdatePerson(ctx context.Context, id int, person Person) (*Person, error) {
	uri := fmt.Sprintf("/persons/%v", id)
	req, err := c.NewRequest(http.MethodPut, uri, nil, person)

	if err != nil {
		return nil, err
	}

	var personResponse *PersonResponse

	_, err = c.Do(ctx, req, &personResponse)

	if err != nil {
		return nil, err
	}

	return personResponse.Data, nil
}

// DeletePerson marks person as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id
func (c *Client) DeletePerson(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/persons/%v", id)
	req, err := c.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return c.Do(ctx, req, nil)
}

// DeleteMultiplePersons marks multiple persons as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons
func (c *Client) DeleteMultiplePersons(ctx context.Context, ids []int) (*Response, error) {
	req, err := c.NewRequest(http.MethodDelete, "/persons", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return c.Do(ctx, req, nil)
}

// // FindByEmail Searches all persons by their Email
// //
// // Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_find
// func (c *Client) FindPersonsByEmail(ctx context.Context, email string) (*FindPersonsResponse, *Response, error) {
// 	req, err := c.NewRequest(http.MethodGet, "/persons/find", struct {
// 		Term          string `url:"term"`
// 		SearchByEmail int    `url:"search_by_email"`
// 	}{
// 		email,
// 		1,
// 	}, nil)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	var personResponse *FindPersonsResponse

// 	_, err = c.Do(ctx, req, &personResponse)

// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return nil, personResponse.Data, nil
// }

//GetPerson returns a person by their id
//
//Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_id
func (c *Client) GetPerson(ctx context.Context, id int) (*Person, error) {
	uri := fmt.Sprintf("/persons/%v", id)
	req, err := c.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	var personResponse *PersonResponse

	_, err = c.Do(ctx, req, &personResponse)

	if err != nil {
		return nil, err
	}

	return personResponse.Data, nil
}

// ListPersons list all persons based on a filter
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons
func (c *Client) ListPersons(ctx context.Context, opt *FilterOptions) ([]*Person, error) {
	req, err := c.NewRequest(http.MethodGet, "/persons", opt, nil)

	if err != nil {
		return nil, err
	}

	var personsResponse *PersonsResponse

	_, err = c.Do(ctx, req, &personsResponse)

	if err != nil {
		return nil, err
	}

	return personsResponse.Data, nil
}
