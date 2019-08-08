package k8

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestCreateNamespaceOperator(t *testing.T) {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)
	kclient, err := createClientset()
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

func TestDeleteNamespaceOperator(t *testing.T) {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)
	kclient, err := createClientset()
	if err != nil {
		panic(err.Error())
	}
	status, err := NewNamespaceController(kclient).DeleteNewNameSpace()
	if err != nil {
		panic(err.Error())
	}
	log.Printf(fmt.Sprintf("Namespace deleted with status %v", status))

	log.Printf("Shutting down...")
}
