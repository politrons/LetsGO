package namespace

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Controller struct {
	kclient *kubernetes.Clientset
}

/**
Factory of the Controller which require the mandatory arguments of the type, which it will use in the rest of the functions below.
*/
func NewPodController(kclient *kubernetes.Clientset) *Controller {
	return &Controller{kclient: kclient}
}

/**
Function to get all pods in the namespace, we just use [List] operator which return a [PodList]
and using the [Item] we receive the array of [Pod]
god bless type system
*/
func (controller *Controller) GetAllPods(namespace string) ([]v1.Pod, error) {
	listOptions := metav1.ListOptions{}
	podList, err := controller.kclient.CoreV1().Pods(namespace).List(listOptions)

	if err != nil {
		return nil, err
	}
	return podList.Items, err
}

func (controller *Controller) CreatePod(namespace string) (*v1.Pod, error) {
	podInfo := controller.createPodInfo(namespace)
	pod, err := controller.kclient.CoreV1().Pods(namespace).Create(podInfo)

	if err != nil {
		return nil, err
	}
	return pod, err
}

func (controller *Controller) DeletePod(namespace string, podName string) (bool, error) {
	var deleteOptions *metav1.DeleteOptions
	err := controller.kclient.CoreV1().Pods(namespace).Delete(podName, deleteOptions)

	if err != nil {
		return false, err
	}
	return true, err
}

func (controller *Controller) createPodInfo(namespace string) *v1.Pod {
	return &v1.Pod{TypeMeta: metav1.TypeMeta{
		Kind:       "pod",
		APIVersion: "rbac.authorization.k8s.io/v2",
	},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("pod-%s", namespace),
			Namespace: namespace,
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "nginx",
					Image: "nginx",
				},
			},
		},
	}
}
