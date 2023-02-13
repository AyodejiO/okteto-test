package app

import (
	"net/http"
	"os"

	"github.com/AyodejiO/okteto/dto"
	"github.com/AyodejiO/okteto/service"
)

type PodHandlers struct {
	service service.PodService
}

func (h PodHandlers) GetPodsCount(w http.ResponseWriter, r *http.Request) {
	namespace := os.Getenv("NAMESPACE")

	if count, err := h.service.GetPodsCount(namespace) ; err != nil {
		writeResponse(w, http.StatusForbidden, false, dto.ErrorResponse{Message: err.Error()})
	} else {
		writeResponse(w, http.StatusOK, true, dto.NewPodCountResponse{Count: count})
	}
}

func (h PodHandlers) GetPods(w http.ResponseWriter, r *http.Request) {
	namespace := os.Getenv("NAMESPACE")

	sort := r.URL.Query().Get("sort")

	sortOrder := r.URL.Query().Get("sortOrder")

	if pods, err := h.service.GetPods(namespace) ; err != nil {
		writeResponse(w, http.StatusForbidden, false, dto.ErrorResponse{Message: err.Error()})
	} else {
		writeResponse(w, http.StatusOK, true, dto.NewPodsResponse{Pods: service.SortPods(pods, sort, sortOrder)})
	}
}