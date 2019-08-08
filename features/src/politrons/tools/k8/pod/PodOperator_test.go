package namespace

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"log"
	"politrons/tools/k8"
	"testing"
)

func TestGetAllPodsOperator(t *testing.T) {
	kclient, _ := getKClient()
	pods, err := NewPodController(kclient).GetAllPods("docker")
	if err != nil {
		panic(err.Error())
	}
	for _, pod := range pods {
		log.Printf(fmt.Sprintf("Pod %s in namespace", pod.Name))
	}
	log.Printf("Shutting down...")
}

func TestCreatePodOperator(t *testing.T) {
	kclient, _ := getKClient()
	namespace, err := NewPodController(kclient).CreatePod("default")
	if err != nil {
		panic(err.Error())
	}
	log.Printf(fmt.Sprintf("New Pod %s created", namespace.Name))
	log.Printf("Shutting down...")
}

func TestDeletePodOperator(t *testing.T) {
	kclient, _ := getKClient()
	status, err := NewPodController(kclient).DeletePod("default", "pod-default")
	if err != nil {
		panic(err.Error())
	}
	log.Printf(fmt.Sprintf("Delete Pod status %v", status))
	log.Printf("Shutting down...")
}

func getKClient() (*kubernetes.Clientset, error) {
	kclient, err := k8.CreateClientset()
	if err != nil {
		panic(err.Error())
	}
	return kclient, err
}
