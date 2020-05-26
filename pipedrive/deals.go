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

type Deal interface {
	APIName() string
}

// Deal represents a Pipedrive deal.
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

func (d BaseDealObject) APIName() string {
	return "deals"
}

func (d BaseDealObject) String() string {
	return Stringify(d)
}

// DealsResponse represents multiple deals response.
type DealsResponse struct {
	Success        bool            `json:"success,omitempty"`
	Data           []*Deal         `json:"data,omitempty"`
	AdditionalData *AdditionalData `json:"additional_data,omitempty"`
}

// DealResponse represents single deal response.
type DealResponse struct {
	Success        bool            `json:"success,omitempty"`
	Data           *Deal           `json:"data,omitempty"`
	AdditionalData *AdditionalData `json:"additional_data,omitempty"`
}

// CreateDeal creates a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals
func (c *Client) CreateDeal(ctx context.Context, deal Deal) (*Deal, error) {
	uri := fmt.Sprintf("/deals")
	req, err := c.NewRequest(http.MethodPost, uri, nil, deal)

	if err != nil {
		return nil, err
	}

	var dealResp *DealResponse

	_, err = c.Do(ctx, req, &dealResp)
	if err != nil {
		return nil, err
	}

	return dealResp.Data, nil
}

// UpdateDeal updates a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id
func (c *Client) UpdateDeal(ctx context.Context, id int, deal Deal) (*Deal, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := c.NewRequest(http.MethodPut, uri, nil, deal)

	if err != nil {
		return nil, err
	}

	var dealResponse *DealResponse

	_, err = c.Do(ctx, req, &dealResponse)
	if err != nil {
		return nil, err
	}

	return dealResponse.Data, nil
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

	_, err = c.Do(ctx, req, nil)
	return err
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

	_, err = c.Do(ctx, req, nil)
	return err
}

// FindDeals by name.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_find
// func (c *Client) FindDeals(ctx context.Context, term string) ([]*Deal, error) {
// 	req, err := c.NewRequest(http.MethodGet, "/deals/find", &SearchOptions{
// 		Term: term,
// 	}, nil)

// 	if err != nil {
// 		return nil, err
// 	}

// 	var dealResp *DealsResponse

// 	_, err = c.Do(ctx, req, &dealResp)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return dealResp.Data, nil
// }

type FilterOptions struct {
	FilterID int    `url:"filter_id"`
	Status   string `url:"status"` // e.g. "all_not_deleted"
}

// ListDeals lists deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals
func (c *Client) ListDeals(ctx context.Context, opt *FilterOptions) ([]*Deal, error) {
	req, err := c.NewRequest(http.MethodGet, "/deals", opt, nil)

	if err != nil {
		return nil, err
	}

	var dealsResp *DealsResponse

	_, err = c.Do(ctx, req, &dealsResp)

	if err != nil {
		return nil, err
	}

	return dealsResp.Data, nil
}

// GetDeal gets a deal by ID
func (c *Client) GetDeal(ctx context.Context, id int) (*Deal, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := c.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	var dealResp *DealResponse

	_, err = c.Do(ctx, req, &dealResp)
	if err != nil {
		return nil, err
	}

	return dealResp.Data, nil
}
