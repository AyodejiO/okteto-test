package service

import (
	"context"
	"sort"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodService interface {
	GetPods(namespace string) ([]v1.Pod, error)
	GetPodsCount(namespace string) (int, error)
}

type DefaultPodService struct {
	clientset kubernetes.Interface
}

func (s DefaultPodService) GetPods(namespace string) ([]v1.Pod, error) {
	pods, err := s.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods.Items, nil
}

func (s DefaultPodService) GetPodsCount(namespace string) (int, error) {
	pods, err := s.clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return 0, err
	}

	return len(pods.Items), nil
}

func SortPods(pods []v1.Pod, sort string, sortOrder string) []v1.Pod {
	var asc bool
	if sortOrder == "desc" {
		asc = false
	} else {
		asc = true // defaults to ascending order
	}

	if sort == "name" {
		sortByName(pods, asc)
	} else if sort == "restart_count" {
		sortByRestartCount(pods, asc)
	} else if sort == "age" {	
		sortByAge(pods, asc)
	}

	return pods
}

func sortByName(pods []v1.Pod, asc bool) {
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

func sortByRestartCount(pods []v1.Pod, asc bool) {
	if asc {
		sort.SliceStable(pods, func(i, j int) bool {
			return pods[i].Status.ContainerStatuses[0].RestartCount < pods[j].Status.ContainerStatuses[0].RestartCount
		})
	} else {
		sort.SliceStable(pods, func(i, j int) bool {
			return pods[i].Status.ContainerStatuses[0].RestartCount > pods[j].Status.ContainerStatuses[0].RestartCount
		})
	}
}

func sortByAge(pods []v1.Pod, asc bool) {
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

func NewPodService(clientset kubernetes.Interface) DefaultPodService {
	return DefaultPodService{clientset}
}