package delivery

import (
    "app/internal/domain"
    "net/http"
)

type Handler struct {
    service domain.Service
}

func (h *Handler) SearchByGlobalID(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) SearchByMode(w http.ResponseWriter, r *http.Request) {

}
