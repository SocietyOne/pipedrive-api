package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Deals

type DealStatus string

const (
	Open    DealStatus = "open"
	Won     DealStatus = "won"
	Lost    DealStatus = "lost"
	Deleted DealStatus = "deleted"
)

//go:generate moq -out mock_deal.go . Deal

// Deal represents a Pipedrive deal.
// Should embed BasePersonObject
type Deal interface {
}

// BaseDealObject represents a basic Pipedrive deal.
type BaseDealObject struct {
	// Unsettable Fields
	ID         int    `json:"id,omitempty"`
	Active     bool   `json:"active,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
	PersonName string `json:"person_name,omitempty"`
	OrgName    string `json:"org_name,omitempty"`
	OwnerName  string `json:"owner_name,omitempty"`

	// Settable Fields
	Title        *string     `json:"title,omitempty"`
	Value        *float64    `json:"value,omitempty"`
	StageID      *int        `json:"stage_id,omitempty"`
	PersonID     interface{} `json:"person_id,omitempty"`
	OrgID        interface{} `json:"org_id"`
	UserID       interface{} `json:"user_id,omitempty"`
	Status       *DealStatus `json:"status,omitempty"`
	LostReason   *string     `json:"lost_reason,omitempty"`
	PipelineID   *int        `json:"pipeline_id,omitempty"`
	StageOrderNr *int        `json:"stage_order_nr,omitempty"`

	// Unused fields
	// Currency   string `json:"currency,omitempty"`
	// AddTime    string `json:"add_time,omitempty"`
	// UpdateTime string `json:"update_time,omitempty"`
	// CreatorUserID   interface{} `json:"creator_user_id,omitempty"`
	// StageChangeTime string      `json:"stage_change_time,omitempty"`
	// Probability      interface{} `json:"probability,omitempty"`
	// NextActivityDate interface{} `json:"next_activity_date,omitempty"`
	// NextActivityTime interface{} `json:"next_activity_time,omitempty"`
	// NextActivityID   interface{} `json:"next_activity_id,omitempty"`
	// LastActivityID   int         `json:"last_activity_id,omitempty"`
	// LastActivityDate string      `json:"last_activity_date,omitempty"`
	// VisibleTo                string      `json:"visible_to,omitempty"`
	// CloseTime                string      `json:"close_time,omitempty"`
	// WonTime                  interface{} `json:"won_time,omitempty"`
	// FirstWonTime             interface{} `json:"first_won_time,omitempty"`
	// LostTime                 string      `json:"lost_time,omitempty"`
	// ProductsCount            int         `json:"products_count,omitempty"`
	// FilesCount               int         `json:"files_count,omitempty"`
	// NotesCount               int         `json:"notes_count,omitempty"`
	// FollowersCount           int         `json:"followers_count,omitempty"`
	// EmailMessagesCount       int         `json:"email_messages_count,omitempty"`
	// ActivitiesCount          int         `json:"activities_count,omitempty"`
	// DoneActivitiesCount      int         `json:"done_activities_count,omitempty"`
	// UndoneActivitiesCount    int         `json:"undone_activities_count,omitempty"`
	// ReferenceActivitiesCount int         `json:"reference_activities_count,omitempty"`
	// ParticipantsCount        int         `json:"participants_count,omitempty"`
	// ExpectedCloseDate        interface{} `json:"expected_close_date,omitempty"`
	// LastIncomingMailTime     interface{} `json:"last_incoming_mail_time,omitempty"`
	// LastOutgoingMailTime     interface{} `json:"last_outgoing_mail_time,omitempty"`
	// NextActivitySubject    interface{} `json:"next_activity_subject,omitempty"`
	// NextActivityType       interface{} `json:"next_activity_type,omitempty"`
	// NextActivityDuration   interface{} `json:"next_activity_duration,omitempty"`
	// NextActivityNote       interface{} `json:"next_activity_note,omitempty"`
	// FormattedValue         string      `json:"formatted_value,omitempty"`
	// RottenTime             interface{} `json:"rotten_time,omitempty"`
	// WeightedValue          float64     `json:"weighted_value,omitempty"`
	// FormattedWeightedValue string      `json:"formatted_weighted_value,omitempty"`
	// CcEmail                string      `json:"cc_email,omitempty"`
	// OrgHidden              bool        `json:"org_hidden,omitempty"`
	// PersonHidden           bool        `json:"person_hidden,omitempty"`
}

// CreateDeal creates a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals
func (c *Client) CreateDeal(ctx context.Context, deal Deal, out ResponseModel) error {

	req, err := c.NewRequest(http.MethodPost, "/deals", nil, deal)
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

// UpdateDeal updates a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id
func (c *Client) UpdateDeal(ctx context.Context, id int, deal Deal, out ResponseModel) error {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := c.NewRequest(http.MethodPut, uri, nil, deal)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, out)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDeal deletes a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id
func (c *Client) DeleteDeal(ctx context.Context, id int) error {
	uri := fmt.Sprintf("/deals/%v", id)
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

// DeleteDeals deletes deals in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals
func (c *Client) DeleteDeals(ctx context.Context, ids []int) error {
	req, err := c.NewRequest(http.MethodDelete, "/deals", &DeleteMultipleOptions{
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

// SearchDealsOptions is used to configure a search request. Term is required
type SearchDealsOptions struct {
	Term           string  `url:"term"`                      // The search term to look for. Minimum 2 characters (or 1 if using exact_match). (REQUIRED)
	Fields         *string `url:"fields,omitempty"`          // A comma-separated string array. The fields to perform the search from. Defaults to all of them.
	ExactMatch     *bool   `url:"status,omitempty"`          // When enabled, only full exact matches against the given term are returned. It is not case sensitive.
	PersonID       *int    `url:"person_id,omitempty"`       // Will filter Deals by the provided Person ID. The upper limit of found Deals associated with the Person is 2000.
	OrganizationID *int    `url:"organization_id,omitempty"` // Will filter Deals by the provided Organization ID. The upper limit of found Deals associated with the Organization is 2000.
	IncludeFields  *string `url:"include_fields,omitempty"`  // Supports including optional fields in the results which are not provided by default.
	Start          *int    `url:"start,omitempty"`           // Pagination start.
	Limit          *int    `url:"limit,omitempty"`           // Items shown per page
}

// SearchDeals searches all deals
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_search
func (c *Client) SearchDeals(ctx context.Context, opt *SearchDealsOptions, out ResponseModel) error {
	req, err := c.NewRequest(http.MethodGet, "/deals/search", opt, nil)
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

// ListDealOptions is used to configure a list deals request. PersonID is required
type ListDealOptions struct {
	Status *string `url:"status,omitempty"` // Only fetch deals with specific status. If omitted, all not deleted deals are fetched
	Start  *int    `url:"start,omitempty"`  // Pagination start
	Limit  *int    `url:"limit,omitempty"`  // Items shown per page
	Sort   *string `url:"sort,omitempty"`   // Field names and sorting mode separated by a comma (field_name_1 ASC, field_name_2 DESC). Only first-level field keys are supported (no nested keys)
}

// ListDeals lists deals belonging to a person
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Persons/get_persons_id_deals
func (c *Client) ListDeals(ctx context.Context, personID int, opt *ListDealOptions, out ResponseModel) error {
	uri := fmt.Sprintf("/persons/%v/deals", personID)
	req, err := c.NewRequest(http.MethodGet, uri, opt, nil)
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

// GetDeal gets a deal by ID
func (c *Client) GetDeal(ctx context.Context, id int, out ResponseModel) error {
	uri := fmt.Sprintf("/deals/%v", id)
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
