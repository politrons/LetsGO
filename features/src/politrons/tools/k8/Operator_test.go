package k8

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"testing"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func TestAddRoleBindingOperator(t *testing.T) {
	// Set logging output to standard console out
	log.SetOutput(os.Stdout)
	stop := make(chan struct{})    // Create channel to receive stop signal
	waitGroup := &sync.WaitGroup{} // Goroutines can add themselves to this to be waited on so that they finish
	clientset, err := createClientset()
	if err != nil {
		panic(err.Error())
	}
	NewNamespaceController(clientset).
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
	clientset, err := createClientset()
	if err != nil {
		panic(err.Error())
	}
	NewNamespaceController(clientset).
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
	clientset, err := createClientset()
	if err != nil {
		panic(err.Error())
	}
	NewNamespaceController(clientset).
		AddDeleteRoleBindingEventHandler().
		Run(stop, waitGroup)

	log.Printf("Shutting down...")

	close(stop)      // Tell goroutines to stop themselves
	waitGroup.Wait() // Wait for all to be stopped
}

/**
Function that create the [Clientset] which is basically the Kubernetes client that contains all
Clients to make API calls to the Kubernetes API.

We use [os.Getenv("HOME")] to get the home path, and in there we use the the current context in kubeconfig [~/.kube/config]

Finally we use [kubernetes.NewForConfig] passing the config to create the [Clientset]
*/
func createClientset() (*kubernetes.Clientset, error) {
	homeDir := os.Getenv("HOME")
	kubeConfigLocation := filepath.Join(homeDir, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigLocation)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
