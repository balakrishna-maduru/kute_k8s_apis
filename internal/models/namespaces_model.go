package models

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// GetNamespaces retrieves a list of namespaces from the Kubernetes cluster
func GetNamespaces(clientset *kubernetes.Clientset) ([]v1.Namespace, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return namespaces.Items, nil
}

// CreateNamespace creates a new namespace in the Kubernetes cluster
func CreateNamespace(clientset *kubernetes.Clientset, name string) error {
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
	}
	_, err := clientset.CoreV1().Namespaces().Create(context.Background(), namespace, metav1.CreateOptions{})
	if err != nil {
		return err
	}
	return nil
}

func DeleteNamespace(clientset *kubernetes.Clientset, namespaceName string) error {
	deleteOptions := metav1.DeleteOptions{}
	err := clientset.CoreV1().Namespaces().Delete(context.Background(), namespaceName, deleteOptions)
	if err != nil {
		return fmt.Errorf("error deleting namespace: %v", err)
	}
	return nil
}
