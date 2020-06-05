package pipedrive

import (
	"fmt"
)

const (
	VisibleToOwnersAndFollowers = 1
	VisibleToWholeCompany       = 3
)

type Pagination struct {
	Start                 int  `json:"start"`
	Limit                 int  `json:"limit"`
	MoreItemsInCollection bool `json:"more_items_in_collection"`
}

type AdditionalData struct {
	User struct {
		Profile struct {
			ID              int         `json:"id"`
			Email           string      `json:"email"`
			Name            string      `json:"name"`
			IsAdmin         bool        `json:"is_admin"`
			DefaultCurrency string      `json:"default_currency"`
			IconURL         interface{} `json:"icon_url"`
			Activated       bool        `json:"activated"`
		} `json:"profile"`
		Locale struct {
			Language        string `json:"language"`
			Country         string `json:"country"`
			Uses12HourClock bool   `json:"uses_12_hour_clock"`
		} `json:"locale"`
		Timezone struct {
			Name   string `json:"name"`
			Offset int    `json:"offset"`
		} `json:"timezone"`
	} `json:"user"`
	MultipleCompanies   bool       `json:"multiple_companies"`
	DefaultCompanyID    int        `json:"default_company_id"`
	CompanyID           int        `json:"company_id"`
	SinceTimestamp      string     `json:"since_timestamp"`
	LastTimestampOnPage string     `json:"last_timestamp_on_page"`
	Pagination          Pagination `json:"pagination"`
}

// ResponseModel is the response model
// Should use the BaseResponse with expected struct as BaseResponse.Data.
type ResponseModel interface {
	Successful() bool
	ErrorString() string
}

// BaseResponse is the base response model
type BaseResponse struct {
	Success   bool        `json:"success,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Error     string      `json:"error,omitempty"`
	ErrorInfo string      `json:"error_info,omitempty"`

	AdditionalData AdditionalData `json:"additional_data,omitempty"`
	RelatedObjects interface{}    `json:"related_objects,omitempty"`
}

func (b *BaseResponse) Successful() bool {
	return b.Success
}

func (b *BaseResponse) ErrorString() string {
	return b.Error
}

type DeleteMultipleOptions struct {
	Ids string `url:"ids,omitempty"`
}

type ErrorFields struct {
	Error     string `json:"error"`
	ErrorInfo string `json:"error_info"`
}

// Type of actions.
type EventAction string

const (
	ACTION_ADDED   EventAction = "added"
	ACTION_UPDATED EventAction = "updated"
	ACTION_MERGED  EventAction = "merged"
	ACTION_DELETED EventAction = "deleted"
	ACTION_ALL     EventAction = "all"
)

// Type of objects.
type EventObject string

const (
	OBJECT_ACTIVITY      EventObject = "activity"
	OBJECT_ACTIVTIY_TYPE EventObject = "activity_type"
	OBJECT_DEAL          EventObject = "deal"
	OBJECT_NOTE          EventObject = "note"
	OBJECT_ORGANIZATION  EventObject = "organization"
	OBJECT_PERSON        EventObject = "person"
	OBJECT_PIPELINE      EventObject = "pipeline"
	OBJECT_PRODUCT       EventObject = "product"
	OBJECT_STAGE         EventObject = "stage"
	OBJECT_USER          EventObject = "user"
	OBJECT_ALL_          EventObject = "*"
)

// Active flags
type ActiveFlag uint8

const (
	ActiveFlagEnabled  ActiveFlag = 1
	ActiveFlagDisabled ActiveFlag = 0
)

// Field types
type FieldType string

const (
	FieldTypeVarchar     FieldType = "varchar"
	FieldTypeVarcharAuto FieldType = "varchar_auto"
	FieldTypeText        FieldType = "text"
	FieldTypeDouble      FieldType = "double"
	FieldTypeMonetary    FieldType = "monetary"
	FieldTypeDate        FieldType = "date"
	FieldTypeSet         FieldType = "set"
	FieldTypeEnum        FieldType = "enum"
	FieldTypeUser        FieldType = "user"
	FieldTypeOrg         FieldType = "org"
	FieldTypePeople      FieldType = "people"
	FieldTypePhone       FieldType = "phone"
	FieldTypeTime        FieldType = "time"
	FieldTypeTimerange   FieldType = "timerange"
	FieldTypeDaterange   FieldType = "daterange"
)

// Visiblity
type VisibleTo uint8

const (
	VisibleToOwnersFollowers VisibleTo = 1
	VisibleToEntireCompany   VisibleTo = 3
)

// Deal probability
type DealProbability uint8

const (
	DealProbabilityEnabled  DealProbability = 1
	DealProbabilityDisabled DealProbability = 0
)

// Search
type SearchOptions struct {
	Term string `url:"term,omitempty"`
}

type OrgID struct {
	// Settable Fields
	ID int `json:"value,omitempty"`

	// Unsettable Fields
	OwnerID     int    `json:"owner_id,omitempty"`
	Name        string `json:"name,omitempty"`
	PeopleCount int    `json:"people_count,omitempty"`
	Address     string `json:"address,omitempty"`
	IsActive    bool   `json:"active_flag,omitempty"`
	CCEmail     string `json:"cc_email,omitempty"`
}

//MarshalJSON is a Marshalling override
func (o *OrgID) MarshalJSON() ([]byte, error) {
	format := fmt.Sprintf("\"%d\"", o.ID)
	return []byte(format), nil
}

type UserID struct {
	// Settable Fields
	ID      int `json:"id,omitempty"`
	IDValue int `json:"value,omitempty"`

	// Unsettable Fields
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	IsActive bool   `json:"active_flag,omitempty"`
}

//MarshalJSON is a Marshalling override
func (u *UserID) MarshalJSON() ([]byte, error) {
	format := fmt.Sprintf("\"%d\"", u.ID)
	return []byte(format), nil
}

type PersonID struct {
	// Settable Fields
	ID int `json:"value,omitempty"`

	// Unsettable Fields
	Name     string   `json:"name,omitempty"`
	Email    []*Email `json:"email,omitempty"`
	Phone    []*Phone `json:"phone,omitempty"`
	IsActive bool     `json:"active_flag,omitempty"`
}

//MarshalJSON is a Marshalling override
func (p *PersonID) MarshalJSON() ([]byte, error) {
	format := fmt.Sprintf("\"%d\"", p.ID)
	return []byte(format), nil
}
