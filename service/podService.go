package service

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type PodService interface {
	GetPods(namespace string) ([]v1.Pod, error)
	GetPodsCount(namespace string) (int, error)
}

type DefaultPodService struct {
	clientset *kubernetes.Clientset
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

func NewPodService(clientset *kubernetes.Clientset) DefaultPodService {
	return DefaultPodService{clientset}
}