package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Event is used by pop to map your events database table to your go code.
type Event struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Content   string    `json:"content" db:"content"`
	AccountID string    `json:"account_id" db:"account_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// {"partnerCallBackType":"CREATED_ACCOUNT","url":"","publicUrl":"","emails":["matthieub+0318@dossierfacile.fr"],"internalPartnersId":["54"],"tenantLightAPIInfos":[{"email":"matthieub+0318@dossierfacile.fr","salary":11000,"tenantSituation":"CDI","guarantor":true,"listFiles":{"tenantFile5":null,"tenantFile4":null,"tenantFile3":null,"tenantFile2":"d594485f-b41a-4726-91c3-7135356eda18.pdf","tenantFile1":"38920521-62a1-4013-8559-cf3c02d8b18c.pdf","guarantorFile0":"b39bccd1-8353-4546-ab8f-a460149a7b54.pdf","guarantorFile1":"076c51b2-2625-4681-847d-d7f5045f26ed.pdf"},"allInternalPartnerId":["54"]}]}

// String is not required by pop and may be deleted
func (e Event) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Events is not required by pop and may be deleted
type Events []Event

// String is not required by pop and may be deleted
func (e Events) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Event) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Content, Name: "Content"},
		&validators.StringIsPresent{Field: e.AccountID, Name: "AccountID"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
