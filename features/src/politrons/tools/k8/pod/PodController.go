package namespace

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

/**
Pods are the smallest deployable units of computing that can be created and managed in Kubernetes

A Pod is a group of one or more containers(docker), with shared storage/network,
and a specification for how to run the containers.

You can run a shell in the Pod using command [kubectl exec -it your-pod-name -- /bin/bash]
*/

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

/**
Function to create a new [Pod] for the creation of the Pod like in other K8s components like RoleBinding, we need
to pass the instance of the type to persist, in this case [Pod]
Then using a namespace name that already exist we use the API [CoreV1().Pods(namespace).Create]
passing the pod.
*/
func (controller *Controller) CreatePod(namespace string) (*v1.Pod, error) {
	podInfo := controller.createPodInfo(namespace)
	pod, err := controller.kclient.CoreV1().Pods(namespace).Create(podInfo)

	if err != nil {
		return nil, err
	}
	return pod, err
}

/**
Function to delete a Pod using the namespace name where is running and the name of the pod.
*/
func (controller *Controller) DeletePod(namespace string, podName string) (bool, error) {
	var deleteOptions *metav1.DeleteOptions
	err := controller.kclient.CoreV1().Pods(namespace).Delete(podName, deleteOptions)

	if err != nil {
		return false, err
	}
	return true, err
}

/**
Factory function to create an instance of type [Pod].
We provide some info of the pod using [TypeMeta], and also detail of the pod info using [ObjectMeta]

Finally and most important we define the [PodSpec] where we describe the array of containers that the Pod
contains. It's an array of [Container] which each has a name and image. The pod it will pull the image from
docker-hub and it will run inside the pod.
*/
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
					Name:  "nginx", //
					Image: "nginx", //Image you want to pull and run
				},
			},
		},
	}
}
