package k8

import (
	"fmt"
	"log"
	"sync"
	"time"

	"k8s.io/api/core/v1"
	rbacV1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
)

// NamespaceController watches the kubernetes api for changes to namespaces and
// creates a RoleBinding for that particular namespace.
type NamespaceController struct {
	namespaceInformer cache.SharedIndexInformer
	kclient           *kubernetes.Clientset
}

/**
Factory function to create the [NewNamespaceController] type passing the [Clientset] and also
 the SharedIndexInformer to be used add handler to handle the events on resource
*/
func NewNamespaceController(kclient *kubernetes.Clientset) *NamespaceController {
	return &NamespaceController{kclient: kclient, namespaceInformer: createNameSpaceInformer(kclient)}
}

//###########################//
//  	  NAMESPACE 		 //
//########################## //

func CreateNewNameSpace(kclient *kubernetes.Clientset) (*v1.Namespace, error) {
	namespaceSpec := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: "politrons-ns",
		},
	}

	namespace, err := kclient.CoreV1().Namespaces().Create(namespaceSpec)
	if err != nil {
		return nil, err
	}
	return namespace, nil
}

func DeleteNewNameSpace(kclient *kubernetes.Clientset) (bool, error) {
	deleteOptions := &metav1.DeleteOptions{}
	err := kclient.CoreV1().Namespaces().Delete("politrons-ns", deleteOptions)
	if err != nil {
		return false, err
	}
	return true, nil
}

//###########################//
//  	 ROLE BINDING 		//
//##########################//

/**
Create informer for watching Namespaces interactions use
We use [NewSharedIndexInformer] which create an instance of an element to watch, that instance require the arguments:.
* [ListWatch] is any object that knows how to perform an initial list and start a watch on a resource
* [Object] Namespace resource type we want to watch.
* [defaultEventHandlerResyncPeriod]
* [Indexer]
*/
func createNameSpaceInformer(kclient *kubernetes.Clientset) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				return kclient.CoreV1().Namespaces().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return kclient.CoreV1().Namespaces().Watch(options)
			},
		},
		&v1.Namespace{},
		10*time.Second, //Time to finish and return execution from Run()
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
}

/**
Add event handler to handle the events on resource defined before in [SharedIndexInformer]
We pass as Handler the type [ResourceEventHandlerFuncs] which allow pass three functions:
[AddFunc]: In the function that we pass as Handler
[UpdateFunc]:
[DeleteFunc]
*/
func (controller *NamespaceController) AddCreateRoleBindingEventHandler() *NamespaceController {
	controller.namespaceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.createRoleBindingByNamespace,
	})
	return controller
}

func (controller *NamespaceController) AddUpdateRoleBindingEventHandler() *NamespaceController {
	controller.namespaceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.updateRoleBindingByNamespace,
	})
	return controller
}

func (controller *NamespaceController) AddDeleteRoleBindingEventHandler() *NamespaceController {
	controller.namespaceInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.deleteRoleBindingByNamespace,
	})
	return controller
}

/**
In this function we receive the namespace as obj, so we transform into [Namespace]
and we use it to invoke another functions to create the [roleBinding]
*/
func (controller *NamespaceController) createRoleBindingByNamespace(obj interface{}) {
	namespaceObj := obj.(*v1.Namespace)
	namespaceName := namespaceObj.Name
	roleBinding := controller.createRoleBindingInfo(namespaceName, "RoleBinding")
	controller.addRoleBindingInNameSpace(namespaceName, roleBinding)
}

/**
In this function we receive the namespace as obj, so we transform into [Namespace]
and we use it to invoke another functions to update the [roleBinding]
*/
func (controller *NamespaceController) updateRoleBindingByNamespace(obj interface{}) {
	namespaceObj := obj.(*v1.Namespace)
	namespaceName := namespaceObj.Name
	roleBinding := controller.createRoleBindingInfo(namespaceName, "RoleBinding2")
	controller.updateRoleBindingInNameSpace(namespaceName, roleBinding)
}

/**
In this function we receive the namespace as obj, so we transform into [Namespace]
and we use it to invoke another functions to delete the [roleBinding]
*/
func (controller *NamespaceController) deleteRoleBindingByNamespace(obj interface{}) {
	namespaceObj := obj.(*v1.Namespace)
	controller.deleteRoleBindingInNameSpace(namespaceObj.Name)
}

//Create the roleBinding to pass later to RoleBindings creation
func (controller *NamespaceController) createRoleBindingInfo(namespaceName string, kind string) *rbacV1.RoleBinding {
	roleBinding := &rbacV1.RoleBinding{
		TypeMeta: metav1.TypeMeta{
			Kind:       kind,
			APIVersion: "rbac.authorization.k8s.io/v2",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("ad-kubernetes-%s", namespaceName),
			Namespace: namespaceName,
		},
		Subjects: []rbacV1.Subject{
			{
				Kind: "Group",
				Name: fmt.Sprintf("ad-kubernetes-%s", namespaceName),
			},
		},
		RoleRef: rbacV1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "edit",
		},
	}
	return roleBinding
}

/**
Function to add a [RoleBinding] in the specific namespace
*/
func (controller *NamespaceController) addRoleBindingInNameSpace(namespaceName string, roleBinding *rbacV1.RoleBinding) {
	_, err := controller.kclient.RbacV1().RoleBindings(namespaceName).Create(roleBinding)
	controller.logResponse("Create", err, namespaceName)
}

/**
Function to update a [RoleBinding] in the specific namespace
*/
func (controller *NamespaceController) updateRoleBindingInNameSpace(namespaceName string, roleBinding *rbacV1.RoleBinding) {
	_, err := controller.kclient.RbacV1().RoleBindings(namespaceName).Update(roleBinding)
	controller.logResponse("Update", err, namespaceName)
}

/**
Function to delete a [RoleBinding] in the specific namespace
*/
func (controller *NamespaceController) deleteRoleBindingInNameSpace(namespaceName string) {
	var options *metav1.DeleteOptions
	err := controller.kclient.RbacV1().RoleBindings(namespaceName).Delete(fmt.Sprintf("ad-kubernetes-%s", namespaceName), options)
	controller.logResponse("Delete", err, namespaceName)
}

func (controller *NamespaceController) logResponse(action string, err error, namespaceName string) {
	if err != nil {
		log.Println(fmt.Sprintf("%s :Failed to Role Binding: %s", action, err.Error()))
	} else {
		log.Println(fmt.Sprintf("%s :AD RoleBinding for Namespace: %s", action, namespaceName))
	}
}

/*
Run starts the process for listening for namespace changes and acting upon those changes.
[controller.namespaceInformer.Run] starts and runs the shared informer, returning after it stops.
The informer will be stopped when stopCh is closed

We define the channel in read mode
<-chan // read only
chan<- // write only
chan   // write/read
*/
func (controller *NamespaceController) Run(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	// When this function completes, mark the go function as done
	defer wg.Done()
	// Increment wait group as we're about to execute a go function
	wg.Add(1)
	// Execute go function
	go controller.namespaceInformer.Run(stopCh)

	for {
		select {
		case stopSignal := <-stopCh:
			log.Println(fmt.Sprintf("Stop signal from : %s", stopSignal))
			return
		case <-time.After(time.Duration(5000 * time.Millisecond)):
			log.Println("Error timeout finishing process")
			return
		}
	}

}
