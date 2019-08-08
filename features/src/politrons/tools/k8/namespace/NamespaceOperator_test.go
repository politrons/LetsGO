package namespace

import (
	"fmt"
	"log"
	"os"
	"politrons/tools/k8"
	"testing"
)

func TestCreateNamespaceOperator(t *testing.T) {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)
	kclient, err := k8.CreateClientset()
	if err != nil {
		panic(err.Error())
	}
	namespace, err := NewNamespaceController(kclient).CreateNewNameSpace()
	if err != nil {
		panic(err.Error())
	}
	log.Printf(fmt.Sprintf("New namespace %s created", namespace.Name))

	log.Printf("Shutting down...")
}

func TestUpdateNamespaceOperator(t *testing.T) {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)
	kclient, err := k8.CreateClientset()
	if err != nil {
		panic(err.Error())
	}
	namespace, err := NewNamespaceController(kclient).UpdateNameSpace()
	if err != nil {
		panic(err.Error())
	}
	log.Printf(fmt.Sprintf("Namespace %s updated", namespace.Name))

	log.Printf("Shutting down...")
}

func TestDeleteNamespaceOperator(t *testing.T) {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)
	kclient, err := k8.CreateClientset()
	if err != nil {
		panic(err.Error())
	}
	status, err := NewNamespaceController(kclient).DeleteNameSpace()
	if err != nil {
		panic(err.Error())
	}
	log.Printf(fmt.Sprintf("Namespace deleted with status %v", status))

	log.Printf("Shutting down...")
}
