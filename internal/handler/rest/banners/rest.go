package tournaments

import (
	"encoding/json"
	"net/http"

	"github.com/ibra-bybuy/go-wsports-events/internal/controller/banners"
	"github.com/ibra-bybuy/go-wsports-events/pkg/model"
)

type Handler struct {
	ctrl *banners.Controller
}

func New(c *banners.Controller) *Handler {
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
	h.getAllBanners(w, req)
}

func (h *Handler) getAllBanners(w http.ResponseWriter, req *http.Request) {
	response := h.ctrl.Get(req.Context())

	json.NewEncoder(w).Encode(model.SuccessResponse{
		Success: true,
		Data:    response,
	})
}
