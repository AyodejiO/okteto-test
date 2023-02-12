package dto

import (
	v1 "k8s.io/api/core/v1"
)

type NewPodCountResponse struct {
	Count int `json:"pods_count"`
}

type NewPodsResponse struct {
	Pods []v1.Pod `json:"pods"`
}