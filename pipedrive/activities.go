package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// Pipedrive API dcos: https://developers.pipedrive.com/docs/api/v1/#!/Deals

//go:generate moq -out mock_activity.go . Activity

// Activity represents a Pipedrive activity.
// Should embed BaseActivityObject
type Activity interface {
}

// BaseActivityObject represents a basic Pipedrive deal.
type BaseActivityObject struct {
	// Unsettable Fields
	ID int `json:"id,omitempty"`

	// Settable Fields
	Subject *string `json:"subject,omitempty"` // Subject of the activity
	// Done *bool `json:"done,omitempty"`
	Type              *string `json:"type,omitempty"`               // Type of the activity. This is in correlation with the key_string parameter of ActivityTypes.
	DueDate           *string `json:"due_date,omitempty"`           // Due date of the activity. Format: YYYY-MM-DD
	DueTime           *string `json:"due_time,omitempty"`           // Due time of the activity in UTC. Format: HH:MM
	Duration          *string `json:"duration,omitempty"`           // Duration of the activity. Format: HH:MM
	UserID            *int    `json:"user_id,omitempty"`            // ID of the user whom the activity will be assigned to. If omitted, the activity will be assigned to the authorized user.
	DealID            *int    `json:"deal_id,omitempty"`            // ID of the deal this activity will be associated with
	PersonID          *int    `json:"person_id,omitempty"`          // ID of the person this activity will be associated with
	OrgID             *int    `json:"org_id,omitempty"`             // ID of the organization this activity will be associated with
	Note              *string `json:"note,omitempty"`               // Note of the activity (HTML format)
	PublicDescription *string `json:"public_description,omitempty"` // Additional details about the activity that will be synced to your external calendar. Unlike the note added to the activity, the description will be publicly visible to any guests added to the activity.
}

// CreateActivity creates an activity .
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Activities/post_activities
func (c *Client) CreateActivity(ctx context.Context, activity Activity, out ResponseModel) error {

	req, err := c.NewRequest(http.MethodPost, "/activities", nil, activity)
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
