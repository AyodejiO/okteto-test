package app

import (
	"net/http"
	"os"
	"sort"

	"github.com/AyodejiO/okteto/dto"
	"github.com/AyodejiO/okteto/service"
	v1 "k8s.io/api/core/v1"
)

type PodHandlers struct {
	service service.PodService
}

func (h PodHandlers) GetPodCount(w http.ResponseWriter, r *http.Request) {
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
		writeResponse(w, http.StatusOK, true, dto.NewPodsResponse{Pods: SortPods(pods, sort, sortOrder)})
	}
}

func SortPods(pods []v1.Pod, sort string, sortOrder string) []v1.Pod {
	var asc bool
	if sortOrder == "desc" {
		asc = false
	} else {
		asc = true // defaults to ascending order
	}

	if sort == "name" {
		SortByName(pods, asc)
	} else if sort == "restart_count" {
		SortByRestartCount(pods, asc)
	} else if sort == "age" {	
		SortByAge(pods, asc)
	}

	return pods
}

func SortByName(pods []v1.Pod, asc bool) {
	if asc {
		sort.SliceStable(pods, func(i, j int) bool {
			return pods[i].ObjectMeta.Name < pods[j].ObjectMeta.Name
		})
	} else {
		sort.SliceStable(pods, func(i, j int) bool {
			return pods[i].ObjectMeta.Name > pods[j].ObjectMeta.Name
		})
	}
}

func SortByRestartCount(pods []v1.Pod, asc bool) {
	if asc {
		sort.SliceStable(pods, func(i, j int) bool {
			return pods[i].Status.ContainerStatuses[0].RestartCount > pods[j].Status.ContainerStatuses[0].RestartCount
		})
	} else {
		sort.SliceStable(pods, func(i, j int) bool {
			return pods[i].Status.ContainerStatuses[0].RestartCount < pods[j].Status.ContainerStatuses[0].RestartCount
		})
	}
}

func SortByAge(pods []v1.Pod, asc bool) {
	if asc {
		sort.SliceStable(pods, func(i, j int) bool {
			return pods[i].ObjectMeta.CreationTimestamp.Before(&pods[j].ObjectMeta.CreationTimestamp)
		})
	} else {
		sort.SliceStable(pods, func(i, j int) bool {
			return pods[j].ObjectMeta.CreationTimestamp.Before(&pods[i].ObjectMeta.CreationTimestamp)
		})
	}
}