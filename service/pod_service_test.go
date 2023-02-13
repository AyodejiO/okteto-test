package service

import (
	"reflect"
	"testing"

	"github.com/AyodejiO/okteto/tests"
	v1 "k8s.io/api/core/v1"
)

func setup() []v1.Pod {
	// create Pods
	pods := make([]v1.Pod, 0)
	pods = append(pods, tests.GeneratePodTemplate("pod1", 2))
	pods = append(pods, tests.GeneratePodTemplate("pod2", 1))
	return pods
}

func TestGetPodsCount(t *testing.T) {
	clientset := tests.GetTestClientset()
	podService := NewPodService(clientset)

	count, err := podService.GetPodsCount("default")
	if err != nil {
		t.Error(err)
	}

	if len(clientset.Actions()) != 1 {
		t.Errorf("Expected 1 action in GetPodsCount got: %v", len(clientset.Actions()))
	}

	if count != 0 {
		t.Errorf("Expected 2 got: %v", count)
	}
}

func TestGetPods(t *testing.T) {
	clientset := tests.GetTestClientset()
	podService := NewPodService(clientset)

	pods, err := podService.GetPods("default")
	if err != nil {
		t.Error(err)
	}

	if len(clientset.Actions()) != 1 {
		t.Errorf("Expected 1 action in GetPods got: %v", len(clientset.Actions()))
	}

	if clientset.Actions()[0].GetVerb() != "list" {
		t.Errorf("Expected 'list' action in GetPods got: %v", clientset.Actions()[0].GetVerb())
	}

	if clientset.Actions()[0].GetResource().Resource != "pods" {
		t.Errorf("Expected 'pods' action in GetPods got: %v", clientset.Actions()[0].GetResource().Resource)
	}

	if reflect.TypeOf(pods).String() != "[]v1.Pod" {
		t.Errorf("Expected '[]v1.Pod' type in GetPods got: %v", reflect.TypeOf(pods).String())
	}
}

func TestSortPodsByName(t *testing.T) {
	pods := setup()	
	sortedPods := SortPods(pods, "name", "desc")

	if sortedPods[0].ObjectMeta.Name != "pod2" {
		t.Errorf("Expected 'pod2' got: %v", sortedPods[0].ObjectMeta.Name)
	}
}

func TestSortPodsByAge(t *testing.T) {
	pods := setup()	
	sortedPods := SortPods(pods, "age", "desc")

	if sortedPods[0].ObjectMeta.Name != "pod2" {
		t.Errorf("Expected 'pod2' got: %v", sortedPods[0].ObjectMeta.Name)
	}
}

func TestSortPodsByRestartCountAsc(t *testing.T) {
	pods := setup()	
	sortedPods := SortPods(pods, "restart_count", "")

	if sortedPods[0].ObjectMeta.Name != "pod2" {
		t.Errorf("Expected 'pod2' got: %v", sortedPods[0].ObjectMeta.Name)
	}
}