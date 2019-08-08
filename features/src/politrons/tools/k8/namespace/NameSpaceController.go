package namespace

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Controller struct {
	kclient *kubernetes.Clientset
}

func NewNamespaceController(kclient *kubernetes.Clientset) *Controller {
	return &Controller{kclient: kclient}
}

//###########################//
//  	  NAMESPACE 		 //
//########################## //

func (controller *Controller) CreateNewNameSpace() (*v1.Namespace, error) {
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

func (controller *Controller) UpdateNameSpace() (*v1.Namespace, error) {
	getOptions := metav1.GetOptions{}

	namespace, e := controller.kclient.CoreV1().Namespaces().Get("politrons-ns", getOptions)
	if e != nil {
		return nil, e
	}
	name, err := controller.kclient.CoreV1().Namespaces().Update(namespace)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func (controller *Controller) DeleteNameSpace() (bool, error) {
	deleteOptions := &metav1.DeleteOptions{}
	err := controller.kclient.CoreV1().Namespaces().Delete("politrons-ns", deleteOptions)
	if err != nil {
		return false, err
	}
	return true, nil
}
