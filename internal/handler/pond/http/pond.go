package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/abdulsalam/delos/internal/constant"
	"github.com/abdulsalam/delos/internal/entity"
	"github.com/abdulsalam/delos/internal/entity/generic"
	"github.com/abdulsalam/delos/internal/handler"
	_middlewareHandler "github.com/abdulsalam/delos/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
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

	pond, err := h.pondUsecase.GetAll(ctx, entity.PondRequestWithPagination{
		PaginationRequest: generic.PaginationRequest{
			Limit:  limitAsInt,
			Offset: offsetAsInt,
		},
	})
	if err != nil {
		handler.ErrorHandlerReturn(err, w)
	}

	// Send to middl.
	responseWriter := w.(*_middlewareHandler.ResponseWriter)
	responseWriter.ResponseData = pond
}

func (h *Handler) GetByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, constant.PondID)

	uuidId, err := uuid.Parse(id)
	if err != nil {
		return
	}

	pond, err := h.pondUsecase.GetByID(ctx, uuidId)
	if err != nil {
		handler.ErrorHandlerReturn(err, w)
	}

	// Send to middl.
	responseWriter := w.(*_middlewareHandler.ResponseWriter)
	responseWriter.ResponseData = pond
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := &entity.PondRequest{}
	if err := render.Bind(r, data); err != nil {
		log.Println(err)
		return
	}

	id, err := h.pondUsecase.Create(ctx, *data)
	if err != nil {
		handler.ErrorHandlerReturn(err, w)
	}

	// Send to middl.
	responseWriter := w.(*_middlewareHandler.ResponseWriter)
	responseWriter.ResponseData = id
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := &entity.PondRequest{}
	if err := render.Bind(r, data); err != nil {
		log.Println(err)
		return
	}

	id := chi.URLParam(r, constant.PondID)
	uuidId, err := uuid.Parse(id)
	if err != nil {
		return
	}

	// Attach to ID.
	data.ID = uuidId
	rows, err := h.pondUsecase.UpdateByID(ctx, *data)
	if err != nil {
		handler.ErrorHandlerReturn(err, w)
	}

	// Send to middl.
	responseWriter := w.(*_middlewareHandler.ResponseWriter)
	responseWriter.ResponseData = rows
}

func (h *Handler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, constant.PondID)

	uuidId, err := uuid.Parse(id)
	if err != nil {
		return
	}

	farm, err := h.pondUsecase.DeleteByID(ctx, uuidId)
	if err != nil {
		handler.ErrorHandlerReturn(err, w)
	}

	// Send to middl.
	responseWriter := w.(*_middlewareHandler.ResponseWriter)
	responseWriter.ResponseData = farm
}

func (h *Handler) Upsert(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	data := &entity.PondRequest{}
	if err := render.Bind(r, data); err != nil {
		log.Println(err)
		return
	}

	id, err := h.pondUsecase.Upsert(ctx, *data)
	if err != nil {
		handler.ErrorHandlerReturn(err, w)
	}

	// Send to middl.
	responseWriter := w.(*_middlewareHandler.ResponseWriter)
	responseWriter.ResponseData = id
}
