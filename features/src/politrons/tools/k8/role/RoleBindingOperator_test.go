package role

import (
	"log"
	"os"
	"politrons/tools/k8"
	"sync"
	"testing"
)

func TestAddRoleBindingOperator(t *testing.T) {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)
	stop := make(chan struct{})    // Create channel to receive stop signal
	waitGroup := &sync.WaitGroup{} // Goroutines can add themselves to this to be waited on so that they finish
	clientset, err := k8.CreateClientset()
	if err != nil {
		panic(err.Error())
	}
	NewRoleBindingController(clientset).
		AddCreateRoleBindingEventHandler().
		Run(stop, waitGroup)

	/*	<-sigs // Wait for signals (this hangs until a signal arrives)
	 */
	log.Printf("Shutting down...")

	close(stop)      // Tell goroutines to stop themselves
	waitGroup.Wait() // Wait for all to be stopped
}

func TestUpdateRoleBindingOperator(t *testing.T) {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)
	stop := make(chan struct{})    // Create channel to receive stop signal
	waitGroup := &sync.WaitGroup{} // Goroutines can add themselves to this to be waited on so that they finish
	clientset, err := k8.CreateClientset()
	if err != nil {
		panic(err.Error())
	}
	NewRoleBindingController(clientset).
		AddUpdateRoleBindingEventHandler().
		Run(stop, waitGroup)

	/*	<-sigs // Wait for signals (this hangs until a signal arrives)
	 */
	log.Printf("Shutting down...")

	close(stop)      // Tell goroutines to stop themselves
	waitGroup.Wait() // Wait for all to be stopped
}

func TestDeleteRoleBindingOperator(t *testing.T) {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)
	stop := make(chan struct{})    // Create channel to receive stop signal
	waitGroup := &sync.WaitGroup{} // Goroutines can add themselves to this to be waited on so that they finish
	clientset, err := k8.CreateClientset()
	if err != nil {
		panic(err.Error())
	}
	NewRoleBindingController(clientset).
		AddDeleteRoleBindingEventHandler().
		Run(stop, waitGroup)

	log.Printf("Shutting down...")

	close(stop)      // Tell goroutines to stop themselves
	waitGroup.Wait() // Wait for all to be stopped
}
