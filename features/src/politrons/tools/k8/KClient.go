package k8

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

/**
client-go is being used by Kubernetes as the offical API client library

Function that create the [Clientset] which is basically the Kubernetes client that contains all
Clients to make API calls to the Kubernetes API.

We use [os.Getenv("HOME")] to get the home path, and in there we use the the current context in kubeconfig [~/.kube/config]

Finally we use [kubernetes.NewForConfig] passing the config to create the [Clientset]
*/
func CreateClientset() (*kubernetes.Clientset, error) {
	homeDir := os.Getenv("HOME")
	kubeConfigLocation := filepath.Join(homeDir, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigLocation)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}
