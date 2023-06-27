package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/abdulsalam/delos/internal/constant"
	"github.com/abdulsalam/delos/internal/entity"
	"github.com/abdulsalam/delos/internal/entity/generic"
	_middlewareHandler "github.com/abdulsalam/delos/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	responseWriter := w.(*_middlewareHandler.ResponseWriter)
	ctx := r.Context()

	query := r.URL.Query()
	limit := query.Get(constant.Limit)
	offset := query.Get(constant.Offset)

	// Convert to int.
	limitAsInt, err := strconv.Atoi(limit)
	if err != nil {
		limitAsInt = constant.DefaultLimit
	}
	offsetAsInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetAsInt = constant.DefaultOffset
	}

	farm, err := h.usecase.GetAll(ctx, entity.FarmRequestWithPagination{
		PaginationRequest: generic.PaginationRequest{
			Limit:  limitAsInt,
			Offset: offsetAsInt,
		},
	})
	if err != nil {
		responseWriter.StatusCode = http.StatusInternalServerError
		return
	}

	// Send to middl.
	responseWriter.ResponseData = farm
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	responseWriter := w.(*_middlewareHandler.ResponseWriter)

	ctx := r.Context()
	id := chi.URLParam(r, constant.FarmID)

	uuidId, err := uuid.Parse(id)
	if err != nil {
		responseWriter.StatusCode = http.StatusBadRequest
		return
	}

	farm, err := h.usecase.GetByID(ctx, uuidId)
	if err != nil {
		responseWriter.StatusCode = http.StatusInternalServerError
		return
	}

	// Send to middl.
	responseWriter.ResponseData = farm
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	responseWriter := w.(*_middlewareHandler.ResponseWriter)

	ctx := r.Context()
	data := &entity.FarmRequest{}
	if err := render.Bind(r, data); err != nil {
		fmt.Println(err)
		responseWriter.StatusCode = http.StatusBadRequest
		return
	}

	id, err := h.usecase.Create(ctx, *data)
	if err != nil {
		responseWriter.StatusCode = http.StatusInternalServerError
		return
	}

	// Send to middl.
	responseWriter.ResponseData = id
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	responseWriter := w.(*_middlewareHandler.ResponseWriter)

	ctx := r.Context()
	data := &entity.FarmRequest{}
	if err := render.Bind(r, data); err != nil {
		responseWriter.StatusCode = http.StatusBadRequest
		return
	}

	id := chi.URLParam(r, constant.FarmID)
	uuidId, err := uuid.Parse(id)
	if err != nil {
		responseWriter.StatusCode = http.StatusBadRequest
		return
	}

	// Attach to ID.
	data.ID = uuidId
	rows, err := h.usecase.UpdateByID(ctx, *data)
	if err != nil {
		responseWriter.StatusCode = http.StatusInternalServerError
		return
	}

	// Send to middl.
	responseWriter.ResponseData = rows
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	responseWriter := w.(*_middlewareHandler.ResponseWriter)

	ctx := r.Context()
	id := chi.URLParam(r, constant.FarmID)

	uuidId, err := uuid.Parse(id)
	if err != nil {
		responseWriter.StatusCode = http.StatusBadRequest
		return
	}

	farm, err := h.usecase.DeleteByID(ctx, uuidId)
	if err != nil {
		responseWriter.StatusCode = http.StatusInternalServerError
		return
	}

	// Send to middl.
	responseWriter.ResponseData = farm
}

func (h *Handler) Upsert(w http.ResponseWriter, r *http.Request) {
	responseWriter := w.(*_middlewareHandler.ResponseWriter)

	ctx := r.Context()
	data := &entity.FarmRequest{}
	if err := render.Bind(r, data); err != nil {
		responseWriter.StatusCode = http.StatusBadRequest
		return
	}

	id, err := h.usecase.Upsert(ctx, *data)
	if err != nil {
		responseWriter.StatusCode = http.StatusInternalServerError
	}

	// Send to middl.
	responseWriter.ResponseData = id
}
