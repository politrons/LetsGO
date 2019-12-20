package namespace

import (
	"fmt"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
	"time"
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

//###################################//
//      operations over Pod API      //
//################################## //

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
	pod = controller.watchAndReturnPodWhenReady(err, namespace, pod)
	if pod.Status.Phase != v1.PodRunning {
		return nil, fmt.Errorf("Pod is unavailable: %v", pod.Status.Phase)
	}
	return pod, err
}

/**
In order to wait for the Pod to be created we can use [Watch] function, which it will ask for the type
[ListOptions] as argument, and it will return a tuple of [watch.Interface] and error.

Watch.Interface it contains the type [ResultChan] which is a channel that is will send you an [watch.Event]
type, having this Event, we are able to extract the Object which in this case, it should be [Pod]
then casting to [Pod] we can check the state of the Pod, and wait until the [status.Phase] is not [PodPending]
*/
func (controller *Controller) watchAndReturnPodWhenReady(err error, namespace string, pod *v1.Pod) *v1.Pod {
	watch, err := controller.kclient.CoreV1().Pods(namespace).Watch(metav1.ListOptions{
		Watch:           true,
		ResourceVersion: pod.ResourceVersion,
	})
	func() {
		for {
			select {
			case event, ok := <-watch.ResultChan():
				if !ok {
					log.Println("Error In the initialization of Pod.")
				}
				pod = event.Object.(*v1.Pod)
				if pod.Status.Phase != v1.PodPending {
					watch.Stop()
					return
				}
			case <-time.After(10 * time.Minute):
				log.Println("Error Pod took too much time to be created.")
				watch.Stop()
				return
			}
		}
	}()
	return pod
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

For this example we will pull an run Cassandra and Nginx server.
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
					Image: "nginx", //Nginx image you want to pull and run
				},
				{
					Name:  "cassandra", //
					Image: "cassandra", //Cassandra image you want to pull and run
					Ports: []v1.ContainerPort{
						{
							Name:          "client-port",
							Protocol:      v1.ProtocolTCP,
							HostIP:        "0.0.0.0",
							ContainerPort: 9042,
						},
						{
							Name:          "thrift",
							Protocol:      v1.ProtocolTCP,
							HostIP:        "0.0.0.0",
							ContainerPort: 9160,
						},
						{
							Name:          "inter-node",
							Protocol:      v1.ProtocolTCP,
							ContainerPort: 7001,
							HostPort:      7000,
						},
					},
				},
			},
		},
	}
}
