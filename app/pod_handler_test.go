package app

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"

	"testing"

	"github.com/AyodejiO/okteto/service"
	"github.com/AyodejiO/okteto/tests"
	v1 "k8s.io/api/core/v1"
)

type PodsExpectedResponse struct {
	Success bool `json:"success"`
	Data  struct {
		Pods   []v1.Pod `json:"pods"`
		Count  int           `json:"pods_count"`
	} `json:"data"`
}

var ph PodHandler 

var pods []v1.Pod

func init() {
	clientset := tests.GetTestClientset()
	podService := service.NewPodService(clientset)
	ph.service = podService

	// create Pods
	pods = make([]v1.Pod, 0)
	pods = append(pods, tests.GeneratePodTemplate("pod1", 2))
	pods = append(pods, tests.GeneratePodTemplate("pod2", 1))

	// add Pods to clientset
	for _, pod := range pods {
		tests.CreateTestPod(clientset, pod)
	}
}

func setup() PodHandler {
	clientset := tests.GetTestClientset()
	podService := service.NewPodService(clientset)
	ph := PodHandler{service: podService}

	// create Pods
	pods = make([]v1.Pod, 0)
	pods = append(pods, tests.GeneratePodTemplate("pod1", 2))
	pods = append(pods, tests.GeneratePodTemplate("pod2", 1))

	// add Pods to clientset
	for _, pod := range pods {
		tests.CreateTestPod(clientset, pod)
	}

	return ph
}

func TestGetPodsCountHandler(t *testing.T) {
	ph = setup()
	req, err := http.NewRequest("GET", "/pods/count", nil)
	if err != nil {
			t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ph.GetPodsCount)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
	}

	var er PodsExpectedResponse
	if err = json.Unmarshal([]byte(rr.Body.Bytes()), &er); err != nil {
		t.Fatal(err)
	}

	if er.Success != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
				er.Success, true)
	}

	if er.Data.Count != 2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
				er.Data.Count, 2)
	}
}

func TestGetPodsHandler(t *testing.T) {
	ph = setup()
	req, err := http.NewRequest("GET", "/pods", nil)
	if err != nil {
			t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ph.GetPods)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
	}

	var er PodsExpectedResponse
	if err = json.Unmarshal([]byte(rr.Body.Bytes()), &er); err != nil {
		t.Fatal(err)
	}

	if er.Success != true {
		t.Errorf("handler returned unexpected body: got %v want %v",
				er.Success, true)
	}

	if len(er.Data.Pods) != 2 {
		t.Errorf("handler returned unexpected body: got %v want %v",
				len(er.Data.Pods), 2)
	}

	if reflect.TypeOf(er.Data.Pods).String() != "[]v1.Pod" {
		t.Errorf("handler returned unexpected body: got %v want %v",
				reflect.TypeOf(er.Data.Pods).String(), "[]v1.Pod")
	}
}