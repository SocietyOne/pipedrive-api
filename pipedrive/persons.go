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

// Email field struct
type Email struct {
	Label   string `json:"label"`
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

// Phone field struct
type Phone struct {
	Label   string `json:"label"`
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

// Person represents a Pipedrive person.
// Should embed BasePersonObject
type Person interface {
}

// BasePersonObject represents a basic pipedrive person
type BasePersonObject struct {
	// Unsettable Fields
	ID int `json:"id,omitempty" force:"id,omitempty"`

	// Settable Fields
	Name      *string  `json:"name,omitempty"`       // Required
	FirstName *string  `json:"first_name,omitempty"` // Optional
	LastName  *string  `json:"last_name,omitempty"`  // Optional
	Phone     []*Phone `json:"phone,omitempty"`
	Email     []*Email `json:"email,omitempty"`
	OrgID     *OrgID   `json:"org_id,omitempty"`

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

// CreatePerson create a new person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/post_persons
func (c *Client) CreatePerson(ctx context.Context, person Person, out ResponseModel) error {

	req, err := c.NewRequest(http.MethodPost, "/persons", nil, person)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, out)
	if err != nil {
		return err
	}
	if !out.Successful() {
		return fmt.Errorf("not successful, error: %v", out.ErrorString())
	}

	return nil
}

// UpdatePerson update a specific person.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/put_persons_id
func (c *Client) UpdatePerson(ctx context.Context, id int, person Person, out ResponseModel) error {

	uri := fmt.Sprintf("/persons/%v", id)
	req, err := c.NewRequest(http.MethodPut, uri, nil, person)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, out)
	if err != nil {
		return err
	}

	return nil
}

// DeletePerson marks person as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons_id
func (c *Client) DeletePerson(ctx context.Context, id int) error {

	uri := fmt.Sprintf("/persons/%v", id)
	req, err := c.NewRequest(http.MethodDelete, uri, nil, nil)
	if err != nil {
		return err
	}

	out := &BaseResponse{}
	_, err = c.Do(ctx, req, out)
	if err != nil {
		return err
	}
	if !out.Successful() {
		return fmt.Errorf("not successful, error: %v", out.ErrorString())
	}
	return nil
}

// DeletePersons marks multiple persons as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/delete_persons
func (c *Client) DeletePersons(ctx context.Context, ids []int) error {
	req, err := c.NewRequest(http.MethodDelete, "/persons", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)
	if err != nil {
		return err
	}

	out := &BaseResponse{}
	_, err = c.Do(ctx, req, out)
	if err != nil {
		return err
	}
	if !out.Successful() {
		return fmt.Errorf("not successful, error: %v", out.ErrorString()) // As of writing this, endpoint may not report errors correctly
	}

	return nil
}

// SearchPersonsOptions is used to configure a search request. Term is required
type SearchPersonsOptions struct {
	Term           string  `url:"term"`                      // The search term to look for. Minimum 2 characters (or 1 if using exact_match). (REQUIRED)
	Fields         *string `url:"fields,omitempty"`          // A comma-separated string array. The fields to perform the search from. Defaults to all of them.
	ExactMatch     *bool   `url:"exact_match,omitempty"`     // When enabled, only full exact matches against the given term are returned. It is not case sensitive.
	OrganizationID *int    `url:"organization_id,omitempty"` // Will filter Deals by the provided Organization ID. The upper limit of found Deals associated with the Organization is 2000.
	IncludeFields  *string `url:"include_fields,omitempty"`  // Supports including optional fields in the results which are not provided by default.
	Start          *int    `url:"start,omitempty"`           // Pagination start.
	Limit          *int    `url:"limit,omitempty"`           // Items shown per page
}

// SearchPersonsResponse is used to model the search person response
type SearchPersonsResponse struct {
	Success   bool        `json:"success,omitempty"`
	Data      PersonItems `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	ErrorInfo string      `json:"error_info,omitempty"`
}

// PersonItems contains a list of PersonItem
type PersonItems struct {
	Items []PersonItem `json:"items,omitempty"`
}

// PersonItem contains a SearchResultPerson
type PersonItem struct {
	Person      SearchResultPerson `json:"item,omitempty"`
	ResultScore float64            `json:"result_score,omitempty"`
}

// SearchResultPerson is the model of a person from the search person response
type SearchResultPerson struct {
	ID    int      `json:"id,omitempty"`
	Type  string   `json:"type,omitempty"`
	Name  string   `json:"name,omitempty"`
	Phone []string `json:"phone,omitempty"`
	Email []string `json:"email,omitempty"`
	//visibleto
	//owner
	//organization
	//customer fields
	//notes
}

// Search

// SearchPersons Searches all persons
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_search
func (c *Client) SearchPersons(ctx context.Context, opt *SearchPersonsOptions) (*SearchPersonsResponse, error) {
	req, err := c.NewRequest(http.MethodGet, "/persons/search", opt, nil)
	if err != nil {
		return nil, err
	}

	out := &SearchPersonsResponse{}
	_, err = c.Do(ctx, req, out)
	if err != nil {
		return nil, err
	}
	if !out.Success {
		return nil, fmt.Errorf("not successful, error: %v", out.Error)
	}

	return out, nil
}

//GetPerson returns a person by their id
//
//Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_id
func (c *Client) GetPerson(ctx context.Context, id int, out ResponseModel) error {
	uri := fmt.Sprintf("/persons/%v", id)
	req, err := c.NewRequest(http.MethodGet, uri, nil, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, out)
	if err != nil {
		return err
	}
	if !out.Successful() {
		return fmt.Errorf("not successful, error: %v", out.ErrorString())
	}

	return nil
}
