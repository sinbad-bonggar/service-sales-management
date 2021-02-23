package handler

import (
	"net/http"
	"strconv"

	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/service"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/application/utilities"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/infrastructure/server/http/request"
	"github.com/Sinbad-B2B-Platform/service-sales-management/pkg/infrastructure/server/http/response"
	"github.com/go-chi/chi"
)

// PermanentJourneyPlanHandler ...
type PermanentJourneyPlanHandler interface {
	GetPermanentJourneyPlan(w http.ResponseWriter, r *http.Request)
	GetAllPermanentJourneyPlan(w http.ResponseWriter, r *http.Request)
}

type permanentJourneyPlanHandler struct {
	service service.PermanentJourneyPlanService
}

// NewPermanentJourneyPlanHandler ...
func NewPermanentJourneyPlanHandler(s service.PermanentJourneyPlanService) PermanentJourneyPlanHandler {

	return &permanentJourneyPlanHandler{
		service: s,
	}
}

// GetPermanentJourneyPlan ...
func (h *permanentJourneyPlanHandler) GetPermanentJourneyPlan(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	permanentJourneyPlanID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utilities.JSONError(w, r, "ERR_USERID", 400, nil)
		return
	}

	resp, err := h.service.Get(uint(permanentJourneyPlanID))
	if err != nil {
		utilities.JSONError(w, r, "ERR_PORTFOLIO_ID", 400, nil)
		return
	}

	utilities.JSONSuccess(w, r, resp, nil)
}

// GetAllPermanentJourneyPlan ...
func (h *permanentJourneyPlanHandler) GetAllPermanentJourneyPlan(w http.ResponseWriter, r *http.Request) {
	sp := request.SearchPaginationFromContext(r)
	lo := utilities.PaginationToOffsetLimit(sp.Page, sp.Length)

	result, count, err := h.service.GetAll(lo.Limit, lo.Offset, sp.Search)

	if err != nil {
		utilities.JSONError(w, r, "ERR_INTERNAL", 400, nil)
		return
	}

	resp := []response.PermanentJourneyPlanResponse{}

	for _, r := range result {
		resp = append(resp, response.MapFromModel(r))
	}

	meta := utilities.ResponsMeta{
		Total: count,
		Limit: sp.Length,
		Skip:  sp.Page,
	}

	utilities.JSONSuccess(w, r, resp, meta)
}
