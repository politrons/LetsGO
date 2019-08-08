package namespace

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"log"
	"politrons/tools/k8"
	"testing"
)

func TestGetAllNamespaceOperator(t *testing.T) {
	kclient, _ := getKClient()
	namespaces, err := NewNamespaceController(kclient).GetAllNamespaces()
	if err != nil {
		panic(err.Error())
	}
	for _, namespace := range namespaces {
		log.Printf(fmt.Sprintf("Namespace %s in cluster", namespace.Name))
	}
	log.Printf("Shutting down...")
}

func TestCreateNamespaceOperator(t *testing.T) {
	kclient, _ := getKClient()
	namespace, err := NewNamespaceController(kclient).CreateNewNameSpace("politrons-ns")
	if err != nil {
		panic(err.Error())
	}
	log.Printf(fmt.Sprintf("New namespace %s created", namespace.Name))
	log.Printf("Shutting down...")
}

func TestUpdateNamespaceOperator(t *testing.T) {
	kclient, _ := getKClient()
	namespace, err := NewNamespaceController(kclient).UpdateNameSpace("politrons-ns")
	if err != nil {
		panic(err.Error())
	}
	log.Printf(fmt.Sprintf("Namespace %s updated", namespace.Name))

	log.Printf("Shutting down...")
}

func TestDeleteNamespaceOperator(t *testing.T) {
	kclient, _ := getKClient()
	status, err := NewNamespaceController(kclient).DeleteNameSpace("politrons-ns")
	if err != nil {
		panic(err.Error())
	}
	log.Printf(fmt.Sprintf("Namespace deleted with status %v", status))

	log.Printf("Shutting down...")
}

func getKClient() (*kubernetes.Clientset, error) {
	kclient, err := k8.CreateClientset()
	if err != nil {
		panic(err.Error())
	}
	return kclient, err
}
