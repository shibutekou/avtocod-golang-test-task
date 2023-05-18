package delivery

import (
    "app/internal/domain"
    "encoding/json"
    "net/http"
    "strconv"
)

type Handler struct {
    service *domain.Service
}

func NewHandler(service *domain.Service) *Handler {
    return &Handler{service: service}
}

func (h *Handler) InitRoutes() *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/id", h.SearchByGlobalID)
    mux.HandleFunc("/mode", h.SearchByMode)

    return mux
}

func (h *Handler) SearchByGlobalID(w http.ResponseWriter, r *http.Request) {
    globalID := r.URL.Query().Get("id")

    globalIDInt, err := strconv.Atoi(globalID)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte(err.Error()))
    }

    parkingInfo, err := h.service.Parking.ByGlobalID(globalIDInt)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(err.Error()))
    }

    resp, err := json.Marshal(parkingInfo)

    w.Header().Set("Content-Type", "application/json")
    w.Write(resp)
}

func (h *Handler) SearchByMode(w http.ResponseWriter, r *http.Request) {

}
