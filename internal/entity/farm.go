package entity

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/abdulsalam/delos/helper"
	"github.com/abdulsalam/delos/internal/entity/generic"
	"github.com/google/uuid"
)

// Base entity of farm.
type Farm struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"` // Should be unique.
	Location     string    `json:"location"`
	Size         float64   `json:"size"`
	Established  time.Time `json:"established"` // Timestamps.
	Technologies []string  `json:"technologies"`
	Employees    int       `json:"employees"`

	// Extra.
	generic.MetaInfo

	// Ponds child.
	Pond []Pond `json:"ponds,omitempty"`
}

type FarmResponseWithPagination struct {
	Farm                       []Farm `json:"farm"`
	generic.PaginationResponse `json:"pagination"`
}

type FarmRequest struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"` // Should be unique.
	Location     string    `json:"location"`
	Size         float64   `json:"size"`
	Established  string    `json:"established"`
	Technologies string    `json:"technologies"`
	Employees    int       `json:"employees"`
}

type FarmRequestWithPagination struct {
	ID uuid.UUID `json:"id"`
	generic.PaginationRequest
}

func (a *FarmRequest) Bind(r *http.Request) error {
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

func (fq *FarmRequest) ToBaseEntity() Farm {
	dateBase, err := helper.ConvertStringToTime(fq.Established)
	if err != nil {
		return Farm{}
	}

	return Farm{
		ID:           fq.ID,
		Name:         fq.Name,
		Slug:         fq.Slug,
		Location:     fq.Location,
		Size:         fq.Size,
		Established:  dateBase,
		Technologies: strings.Split(fq.Technologies, ","),
		Employees:    fq.Employees,
	}
}
