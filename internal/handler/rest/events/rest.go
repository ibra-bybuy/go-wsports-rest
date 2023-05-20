package events

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ibra-bybuy/go-wsports-events/internal/controller/events"
	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
	"github.com/ibra-bybuy/go-wsports-events/pkg/utils"
)

type Handler struct {
	ctrl *events.Controller
}

func New(c *events.Controller) *Handler {
	return &Handler{c}
}

func (h *Handler) Handle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch req.Method {

	case http.MethodGet:
		h.get(w, req)
	default:
		w.WriteHeader(http.StatusBadGateway)
	}
}

func (h *Handler) get(w http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")

	if id != "" {
		h.getByID(w, req, id)
		return
	}

	limit, err := strconv.Atoi(req.URL.Query().Get("limit"))

	if err != nil || limit == 0 {
		utils.JSONError(w, model.ErrorResponse{
			Message: "invalid limit",
		}, 400)
		return
	}

	page, err := strconv.Atoi(req.URL.Query().Get("page"))

	if err != nil || page == 0 {
		utils.JSONError(w, model.ErrorResponse{
			Message: "invalid page",
		}, 400)
		return
	}

	sport := req.URL.Query().Get("sport")

	if sport != "" {
		h.getBySport(w, req, sport, limit, page)
		return
	}

	query := req.URL.Query().Get("query")

	if query != "" {
		h.getByQuery(w, req, query, limit, page)
		return
	}

	tournament := req.URL.Query().Get("tournament")

	if tournament != "" {
		h.getByTournament(w, req, tournament, limit, page)
		return
	}

	h.getBySport(w, req, "", limit, page)
}

func (h *Handler) getBySport(w http.ResponseWriter, req *http.Request, sport string, limit int, page int) {
	response := h.ctrl.GetBySport(req.Context(), sport, limit, page)

	json.NewEncoder(w).Encode(model.SuccessResponse{
		Success: true,
		Data:    response,
	})
}

func (h *Handler) getByQuery(w http.ResponseWriter, req *http.Request, query string, limit int, page int) {
	response := h.ctrl.GetByQuery(req.Context(), query, limit, page)

	json.NewEncoder(w).Encode(model.SuccessResponse{
		Success: true,
		Data:    response,
	})
}

func (h *Handler) getByTournament(w http.ResponseWriter, req *http.Request, tournament string, limit int, page int) {
	response := h.ctrl.GetByTournament(req.Context(), tournament, limit, page)

	json.NewEncoder(w).Encode(model.SuccessResponse{
		Success: true,
		Data:    response,
	})
}

func (h *Handler) getByID(w http.ResponseWriter, req *http.Request, id string) {
	response, err := h.ctrl.GetByID(req.Context(), id)
	if err != nil {
		utils.JSONError(w, model.ErrorResponse{
			Message: err.Error(),
		}, http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(model.SuccessResponse{
		Success: true,
		Data:    response,
	})
}
