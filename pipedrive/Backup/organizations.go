package pipedrive

import (
	"context"
	"fmt"
	"net/http"
)

// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations

// Organization represents a Pipedrive organization.

type Organization interface {
	APIName() string
}

type BaseOrganizationObject struct {
	// Unsettable Fields
	ID         int    `json:"id"`
	AddTime    string `json:"add_time"`
	UpdateTime string `json:"update_time"`

	// Settable Fields
	Name string `json:"name"`

	// Unused Fields
	// CompanyID int `json:"company_id"`
	// OwnerID   struct {
	// 	ID         int    `json:"id"`
	// 	Name       string `json:"name"`
	// 	Email      string `json:"email"`
	// 	HasPic     bool   `json:"has_pic"`
	// 	PicHash    string `json:"pic_hash"`
	// 	ActiveFlag bool   `json:"active_flag"`
	// 	Value      int    `json:"value"`
	// } `json:"owner_id"`
	// OpenDealsCount                  int         `json:"open_deals_count"`
	// RelatedOpenDealsCount           int         `json:"related_open_deals_count"`
	// ClosedDealsCount                int         `json:"closed_deals_count"`
	// RelatedClosedDealsCount         int         `json:"related_closed_deals_count"`
	// EmailMessagesCount              int         `json:"email_messages_count"`
	// PeopleCount                     int         `json:"people_count"`
	// ActivitiesCount                 int         `json:"activities_count"`
	// DoneActivitiesCount             int         `json:"done_activities_count"`
	// UndoneActivitiesCount           int         `json:"undone_activities_count"`
	// ReferenceActivitiesCount        int         `json:"reference_activities_count"`
	// FilesCount                      int         `json:"files_count"`
	// NotesCount                      int         `json:"notes_count"`
	// FollowersCount                  int         `json:"followers_count"`
	// WonDealsCount                   int         `json:"won_deals_count"`
	// RelatedWonDealsCount            int         `json:"related_won_deals_count"`
	// LostDealsCount                  int         `json:"lost_deals_count"`
	// RelatedLostDealsCount           int         `json:"related_lost_deals_count"`
	// ActiveFlag                      bool        `json:"active_flag"`
	// CategoryID                      interface{} `json:"category_id"`
	// PictureID                       interface{} `json:"picture_id"`
	// CountryCode                     interface{} `json:"country_code"`
	// FirstChar                       string      `json:"first_char"`
	// VisibleTo                       string      `json:"visible_to"`
	// NextActivityDate                string      `json:"next_activity_date"`
	// NextActivityTime                interface{} `json:"next_activity_time"`
	// NextActivityID                  int         `json:"next_activity_id"`
	// LastActivityID                  int         `json:"last_activity_id"`
	// LastActivityDate                string      `json:"last_activity_date"`
	// TimelineLastActivityTime        interface{} `json:"timeline_last_activity_time"`
	// TimelineLastActivityTimeByOwner interface{} `json:"timeline_last_activity_time_by_owner"`
	// Address                         string      `json:"address"`
	// AddressSubpremise               string      `json:"address_subpremise"`
	// AddressStreetNumber             string      `json:"address_street_number"`
	// AddressRoute                    string      `json:"address_route"`
	// AddressSublocality              string      `json:"address_sublocality"`
	// AddressLocality                 string      `json:"address_locality"`
	// AddressAdminAreaLevel1          string      `json:"address_admin_area_level_1"`
	// AddressAdminAreaLevel2          string      `json:"address_admin_area_level_2"`
	// AddressCountry                  string      `json:"address_country"`
	// AddressPostalCode               string      `json:"address_postal_code"`
	// AddressFormattedAddress         string      `json:"address_formatted_address"`
	// OwnerName                       string      `json:"owner_name"`
	// CcEmail                         string      `json:"cc_email"`
}

func (o BaseOrganizationObject) APIName() string {
	return "organizations"
}

func (o BaseOrganizationObject) String() string {
	return Stringify(o)
}

// OrganizationsResponse represents multiple organizations response.
type OrganizationsResponse struct {
	Success        bool            `json:"success"`
	Data           []*Organization `json:"data"`
	AdditionalData AdditionalData  `json:"additional_data,omitempty"`
}

// OrganizationResponse represents single organization response.
type OrganizationResponse struct {
	Success        bool           `json:"success"`
	Data           *Organization  `json:"data"`
	AdditionalData AdditionalData `json:"additional_data,omitempty"`
}

// CreateOrganization a new organizations.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/post_organizations
func (c *Client) CreateOrganization(ctx context.Context, organization *Organization) (*Organization, error) {
	req, err := c.NewRequest(http.MethodPost, "/organizations", nil, organization)

	if err != nil {
		return nil, err
	}

	var organizationResponse *OrganizationResponse

	_, err = c.Do(ctx, req, &organizationResponse)

	if err != nil {
		return nil, err
	}

	return organizationResponse.Data, nil
}

// UpdateOrganization update a specific organization.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/put_persons_id
func (c *Client) UpdateOrganization(ctx context.Context, id int, organization *Organization) (*Organization, error) {
	uri := fmt.Sprintf("/organizations/%v", id)
	req, err := c.NewRequest(http.MethodPut, uri, nil, organization)

	if err != nil {
		return nil, err
	}

	var organizationResponse *OrganizationResponse
	_, err = c.Do(ctx, req, &organizationResponse)

	if err != nil {
		return nil, err
	}

	return organizationResponse.Data, nil
}

// DeleteOrganization marks an organization as deleted.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/delete_organizations_id
func (c *Client) DeleteOrganization(ctx context.Context, id int) (*Response, error) {
	uri := fmt.Sprintf("/organizations/%v", id)
	req, err := c.NewRequest(http.MethodDelete, uri, nil, nil)
	if err != nil {
		return nil, err
	}

	_, err = c.Do(ctx, req, nil)
	return nil, err
}

// DeleteOrganizations deletes multiple organizations in bulk.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/delete_organizations
func (c *Client) DeleteOrganizations(ctx context.Context, ids []int) error {
	req, err := c.NewRequest(http.MethodDelete, "/organizations", &DeleteMultipleOptions{
		Ids: arrayToString(ids, ","),
	}, nil)
	if err != nil {
		return err
	}

	_, err = c.Do(ctx, req, nil)
	return err
}

// ListOrganizations list all organizations.
//
// Pipedrive API docs: https://developers.pipedrive.com/docs/api/v1/#!/Organizations/get_organizations
func (c *Client) ListOrganizations(ctx context.Context) ([]*Organization, error) {
	req, err := c.NewRequest(http.MethodGet, "/organizations", nil, nil)

	if err != nil {
		return nil, err
	}

	var organizationResponse *OrganizationsResponse

	_, err = c.Do(ctx, req, &organizationResponse)

	if err != nil {
		return nil, err
	}

	return organizationResponse.Data, nil
}
