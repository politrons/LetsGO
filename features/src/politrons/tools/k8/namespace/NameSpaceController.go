package namespace

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Controller struct {
	kclient *kubernetes.Clientset
}

/**
Factory of the Controller which require the mandatory arguments of the type, which it will use in the rest of the functions below.
*/
func NewNamespaceController(kclient *kubernetes.Clientset) *Controller {
	return &Controller{kclient: kclient}
}

//###################################//
//      operations over namespace    //
//################################## //
/**
Using the client API [controller.kclient.CoreV1().Namespaces()] we can invoke the different actions to
[List, Get, Create, Update, Delete] Namespaces.
*/

/**
Function to get all namespaces in the cluster, we just use [List] operator which return a [NamespaceList]
god bless type system
*/
func (controller *Controller) GetAllNamespaces() ([]v1.Namespace, error) {
	listOptions := metav1.ListOptions{}
	namespaceList, e := controller.kclient.CoreV1().Namespaces().List(listOptions)
	if e != nil {
		return nil, e
	}
	return namespaceList.Items, nil
}

//Function to create a namespace passing a [Namespace] type with some data
func (controller *Controller) CreateNewNameSpace(name string) (*v1.Namespace, error) {
	namespaceSpec := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	namespace, err := controller.kclient.CoreV1().Namespaces().Create(namespaceSpec)
	if err != nil {
		return nil, err
	}
	return namespace, nil
}

/**
Function to Update a namespace recovering first the previous namespace using [Get] and
passing the namespace name
*/
func (controller *Controller) UpdateNameSpace(name string) (*v1.Namespace, error) {
	getOptions := metav1.GetOptions{}
	namespace, e := controller.kclient.CoreV1().Namespaces().Get(name, getOptions)
	if e != nil {
		return nil, e
	}
	updatedNamespace, err := controller.kclient.CoreV1().Namespaces().Update(namespace)
	if err != nil {
		return nil, err
	}
	return updatedNamespace, nil
}

//Function to delete a [namespace] passing the name of the namespace and also a [DeleteOptions]
func (controller *Controller) DeleteNameSpace(name string) (bool, error) {
	deleteOptions := &metav1.DeleteOptions{}
	err := controller.kclient.CoreV1().Namespaces().Delete(name, deleteOptions)
	if err != nil {
		return false, err
	}
	return true, nil
}
