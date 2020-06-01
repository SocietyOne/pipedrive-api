package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes

// Note represents a Pipedrive note.
type Note interface {
}

// BaseNoteObject represents a basic pipedrive note
type BaseNoteObject struct {
	// Unsettable Fields
	ID int `json:"id,omitempty"`

	// Settable Fields
	UserID   *int    `json:"user_id,omitempty"`
	DealID   *int    `json:"deal_id,omitempty"`
	PersonID *int    `json:"person_id,omitempty"`
	OrgID    *int    `json:"org_id,omitempty"`
	Content  *string `json:"content,omitempty"`

	// Unused Fields
	// AddTime    Timestamp `json:"add_time,omitempty"`
	// UpdateTime Timestamp `json:"update_time,omitempty"`
	// ActiveFlag               bool      `json:"active_flag,omitempty"`
	// PinnedToDealFlag         bool      `json:"pinned_to_deal_flag,omitempty"`
	// PinnedToPersonFlag       bool      `json:"pinned_to_person_flag,omitempty"`
	// PinnedToOrganizationFlag bool      `json:"pinned_to_organization_flag,omitempty"`
	// LastUpdateUserID         int       `json:"last_update_user_id,omitempty"`
}

// CreateNote creates a note.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes_id
func (c *Client) CreateNote(ctx context.Context, note Note, out ResponseModel) error {

	req, err := c.NewRequest(http.MethodPost, "/notes", nil, note)
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

// UpdateNote updates a specific note.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/put_notes_id
func (c *Client) UpdateNote(ctx context.Context, id int, note Note, out ResponseModel) error {

	uri := fmt.Sprintf("/notes/%v", id)
	req, err := c.NewRequest(http.MethodPut, uri, nil, note)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, out)
	if err != nil {
		return err
	}

	return nil
}

// DeleteNote marks note as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/delete_notes_id
func (c *Client) DeleteNote(ctx context.Context, id int) error {

	uri := fmt.Sprintf("/notes/%v", id)
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

// GetNote returns a specific note by id.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Notes/get_notes_id
func (c *Client) GetNote(ctx context.Context, id int, out ResponseModel) error {

	uri := fmt.Sprintf("/notes/%v", id)
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
