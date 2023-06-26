package entity

import (
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/abdulsalam/delos/internal/entity/generic"
)

type Pond struct {
	ID          uuid.UUID `json:"id"`
	FarmID      uuid.UUID `json:"farm_id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"` // Should be unique.
	Size        float64   `json:"size"`
	WaterSource string    `json:"water_source"`

	// Extra.
	generic.MetaInfo
	*Farm `json:"farm,omitempty"`
}

type PondRequest struct {
	ID          uuid.UUID `json:"id"`
	FarmID      uuid.UUID `json:"farm_id"`
	Name        string    `json:"name"`
	Slug        string    `json:"slug"` // Should be unique.
	Size        float64   `json:"size"`
	WaterSource string    `json:"water_source"`
}

type PondRequestWithPagination struct {
	ID uuid.UUID `json:"id"`
	generic.PaginationRequest
}

type PondResponseWithPagination struct {
	Pond                       []Pond `json:"pond"`
	generic.PaginationResponse `json:"pagination"`
}

func (fq *PondRequest) ToBaseEntity() Pond {
	return Pond{
		ID:          fq.ID,
		FarmID:      fq.FarmID,
		Name:        fq.Name,
		Slug:        fq.Slug,
		Size:        fq.Size,
		WaterSource: fq.WaterSource,
	}
}

func (a *PondRequest) Bind(r *http.Request) error {
	// a.Article is nil if no Article fields are sent in the request. Return an
	// error to avoid a nil pointer dereference.
	if a.Name == "" {
		return errors.New("missing required name fields.")
	}

	// a.User is nil if no Userpayload fields are sent in the request. In this app
	// this won't cause a panic, but checks in this Bind method may be required if
	// a.User or further nested fields like a.User.Name are accessed elsewhere.

	// just a post-process after a decode..
	return nil
}
