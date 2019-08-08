package namespace

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespaceController struct {
	kclient *kubernetes.Clientset
}

func NewNamespaceController(kclient *kubernetes.Clientset) *NamespaceController {
	return &NamespaceController{kclient: kclient}
}

//###########################//
//  	  NAMESPACE 		 //
//########################## //

func (controller *NamespaceController) CreateNewNameSpace() (*v1.Namespace, error) {
	namespaceSpec := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "politrons-ns",
		},
	}

	namespace, err := controller.kclient.CoreV1().Namespaces().Create(namespaceSpec)
	if err != nil {
		return nil, err
	}
	return namespace, nil
}

func (controller *NamespaceController) DeleteNewNameSpace() (bool, error) {
	deleteOptions := &metav1.DeleteOptions{}
	err := controller.kclient.CoreV1().Namespaces().Delete("politrons-ns", deleteOptions)
	if err != nil {
		return false, err
	}
	return true, nil
}
