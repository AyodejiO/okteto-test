package app

import (
	"net/http"
	"os"

	"github.com/AyodejiO/okteto/dto"
	"github.com/AyodejiO/okteto/service"
	v1 "k8s.io/api/core/v1"
)

type PodService interface {
	GetPods(namespace string) ([]v1.Pod, error)
	GetPodsCount(namespace string) (int, error)
}

type PodHandler struct {
	service PodService
}

func (h PodHandler) GetPodsCount(w http.ResponseWriter, r *http.Request) {
	namespace := os.Getenv("NAMESPACE")

	if count, err := h.service.GetPodsCount(namespace) ; err != nil {
		writeResponse(w, http.StatusForbidden, false, dto.ErrorResponse{Message: err.Error()})
	} else {
		writeResponse(w, http.StatusOK, true, dto.NewPodCountResponse{Count: count})
	}
}

func (h PodHandler) GetPods(w http.ResponseWriter, r *http.Request) {
	namespace := os.Getenv("NAMESPACE")

	sort := r.URL.Query().Get("sort")

	sortOrder := r.URL.Query().Get("sortOrder")

	if pods, err := h.service.GetPods(namespace) ; err != nil {
		writeResponse(w, http.StatusForbidden, false, dto.ErrorResponse{Message: err.Error()})
	} else {
		writeResponse(w, http.StatusOK, true, dto.NewPodsResponse{Pods: service.SortPods(pods, sort, sortOrder)})
	}
}