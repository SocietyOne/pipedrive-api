package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// DealService handles deals related
// methods of the Pipedrive API.
//
// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Deals
type DealService service

// Deal represents a Pipedrive deal.
type Deal struct {
	ID                       int         `json:"id,omitempty"`
	CreatorUserID            interface{} `json:"creator_user_id,omitempty"`
	UserID                   interface{} `json:"user_id,omitempty"`
	PersonID                 interface{} `json:"person_id,omitempty"`
	OrgID                    interface{} `json:"org_id"`
	StageID                  int         `json:"stage_id,omitempty"`
	Title                    string      `json:"title,omitempty"`
	Value                    float64     `json:"value,omitempty"`
	Currency                 string      `json:"currency,omitempty"`
	AddTime                  string      `json:"add_time,omitempty"`
	UpdateTime               string      `json:"update_time,omitempty"`
	StageChangeTime          string      `json:"stage_change_time,omitempty"`
	Active                   bool        `json:"active,omitempty"`
	Deleted                  bool        `json:"deleted,omitempty"`
	Status                   string      `json:"status,omitempty"`
	Probability              interface{} `json:"probability,omitempty"`
	NextActivityDate         interface{} `json:"next_activity_date,omitempty"`
	NextActivityTime         interface{} `json:"next_activity_time,omitempty"`
	NextActivityID           interface{} `json:"next_activity_id,omitempty"`
	LastActivityID           int         `json:"last_activity_id,omitempty"`
	LastActivityDate         string      `json:"last_activity_date,omitempty"`
	LostReason               string      `json:"lost_reason,omitempty"`
	VisibleTo                string      `json:"visible_to,omitempty"`
	CloseTime                string      `json:"close_time,omitempty"`
	PipelineID               int         `json:"pipeline_id,omitempty"`
	WonTime                  interface{} `json:"won_time,omitempty"`
	FirstWonTime             interface{} `json:"first_won_time,omitempty"`
	LostTime                 string      `json:"lost_time,omitempty"`
	ProductsCount            int         `json:"products_count,omitempty"`
	FilesCount               int         `json:"files_count,omitempty"`
	NotesCount               int         `json:"notes_count,omitempty"`
	FollowersCount           int         `json:"followers_count,omitempty"`
	EmailMessagesCount       int         `json:"email_messages_count,omitempty"`
	ActivitiesCount          int         `json:"activities_count,omitempty"`
	DoneActivitiesCount      int         `json:"done_activities_count,omitempty"`
	UndoneActivitiesCount    int         `json:"undone_activities_count,omitempty"`
	ReferenceActivitiesCount int         `json:"reference_activities_count,omitempty"`
	ParticipantsCount        int         `json:"participants_count,omitempty"`
	ExpectedCloseDate        interface{} `json:"expected_close_date,omitempty"`
	LastIncomingMailTime     interface{} `json:"last_incoming_mail_time,omitempty"`
	LastOutgoingMailTime     interface{} `json:"last_outgoing_mail_time,omitempty"`
	StageOrderNr             int         `json:"stage_order_nr,omitempty"`
	PersonName               string      `json:"person_name,omitempty"`
	OrgName                  string      `json:"org_name,omitempty"`
	NextActivitySubject      interface{} `json:"next_activity_subject,omitempty"`
	NextActivityType         interface{} `json:"next_activity_type,omitempty"`
	NextActivityDuration     interface{} `json:"next_activity_duration,omitempty"`
	NextActivityNote         interface{} `json:"next_activity_note,omitempty"`
	FormattedValue           string      `json:"formatted_value,omitempty"`
	RottenTime               interface{} `json:"rotten_time,omitempty"`
	WeightedValue            int         `json:"weighted_value,omitempty"`
	FormattedWeightedValue   string      `json:"formatted_weighted_value,omitempty"`
	OwnerName                string      `json:"owner_name,omitempty"`
	CcEmail                  string      `json:"cc_email,omitempty"`
	OrgHidden                bool        `json:"org_hidden,omitempty"`
	PersonHidden             bool        `json:"person_hidden,omitempty"`
}

func (d Deal) String() string {
	return Stringify(d)
}

// DealsResponse represents multiple deals response.
type DealsResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           []Deal         `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// DealResponse represents single deal response.
type DealResponse struct {
	Success        bool           `json:"success,omitempty"`
	Data           Deal           `json:"data,omitempty"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// ListUpdates about a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_id_flow
func (s *DealService) ListUpdates(ctx context.Context, id int) (*DataResponse, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/flow", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DataResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Find deals by name.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals_find
func (s *DealService) Find(ctx context.Context, term string) (*DataResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/deals/find", &SearchOptions{
		Term: term,
	}, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DataResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// List all deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/get_deals
func (s *DealService) List(ctx context.Context) (*DataResponse, *Response, error) {
	req, err := s.client.NewRequest(http.MethodGet, "/deals", nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DataResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// Duplicate a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals_id_duplicate
func (s *DealService) Duplicate(ctx context.Context, id int) (*DataResponse, *Response, error) {
	uri := fmt.Sprintf("/deals/%v/duplicate", id)
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DataResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DealsMergeOptions specifices the optional parameters to the
// DealService.Merge method.
type DealsMergeOptions struct {
	MergeWithID uint `url:"merge_with_id,omitempty"`
}

// Merge two deals.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id_merge
func (s *DealService) Merge(ctx context.Context, id int, opt *DealsMergeOptions) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/merge", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, opt)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DealsUpdateOptions specifices the optional parameters to the
// DealService.Update method.
type DealsUpdateOptions struct {
	Title          string `url:"title,omitempty"`
	Value          string `url:"value,omitempty"`
	Currency       string `url:"currency,omitempty"`
	UserID         uint   `url:"user_id,omitempty"`
	PersonID       uint   `url:"person_id,omitempty"`
	OrganizationID uint   `url:"org_id,omitempty"`
	StageID        uint   `url:"stage_id,omitempty"`
	Status         string `url:"status,omitempty"`
	LostReason     string `url:"lost_reason,omitempty"`
	VisibleTo      uint   `url:"visible_to,omitempty"`
}

// Update a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/put_deals_id
func (s *DealService) Update(ctx context.Context, id int, deal interface{}) (*DataResponse, *Response, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := s.client.NewRequest(http.MethodPut, uri, nil, deal)

	if err != nil {
		return nil, nil, err
	}

	var record *DataResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

// DeleteFollower of a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_followers_follower_id
func (s *DealService) DeleteFollower(ctx context.Context, id int, followerID int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/followers/%v", id, followerID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteMultiple deletes deals in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals
func (s *DealService) DeleteMultiple(ctx context.Context, ids []int) (*Response, error) {
	req, err := s.client.NewRequest(http.MethodDelete, "/deals", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteParticipant deletes participant in a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_participants_deal_participant_id
func (s *DealService) DeleteParticipant(ctx context.Context, dealID int, participantID int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/participants/%v", dealID, participantID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete a deal.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id
func (s *DealService) Delete(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// DeleteAttachedProduct deletes attached product.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/delete_deals_id_products_product_attachment_id
func (s *DealService) DeleteAttachedProduct(ctx context.Context, dealID int, productAttachmentID int) (*Response, error) {
	uri := fmt.Sprintf("/deals/%v/products/%v", dealID, productAttachmentID)
	req, err := s.client.NewRequest(http.MethodDelete, uri, nil, nil)

	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Create a deal.
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Deals/post_deals
func (s *DealService) Create(ctx context.Context, deal interface{}) (*DataResponse, *Response, error) {
	uri := fmt.Sprintf("/deals")
	req, err := s.client.NewRequest(http.MethodPost, uri, nil, deal)

	if err != nil {
		return nil, nil, err
	}

	var record *DataResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}

func (s *DealService) Get(ctx context.Context, id int) (*DataResponse, *Response, error) {
	uri := fmt.Sprintf("/deals/%v", id)
	req, err := s.client.NewRequest(http.MethodGet, uri, nil, nil)

	if err != nil {
		return nil, nil, err
	}

	var record *DataResponse

	resp, err := s.client.Do(ctx, req, &record)

	if err != nil {
		return nil, resp, err
	}

	return record, resp, nil
}
