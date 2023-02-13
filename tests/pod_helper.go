package tests

import (
	"context"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func GetTestClientset() *fake.Clientset {
	clientset := fake.NewSimpleClientset()
	return clientset
}

func GeneratePodTemplate(name string, restart_count int) v1.Pod {
	if name == "" {
		name = "test-pod"
	}
	return v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			CreationTimestamp: metav1.Time{
				Time: metav1.Now().Time,
			},
		},
		Status: v1.PodStatus{
			Phase: v1.PodRunning,
			ContainerStatuses: []v1.ContainerStatus{
				{	
					Name: name,
					RestartCount: int32(restart_count),
				},
				{	
					Name: name,
					RestartCount: int32(restart_count),
				},
			},
		},
	}
}

func CreateTestPod(clientset *fake.Clientset, pod v1.Pod) {
	clientset.CoreV1().Pods("test-namespace").Create(context.TODO(), &pod, metav1.CreateOptions{})
}