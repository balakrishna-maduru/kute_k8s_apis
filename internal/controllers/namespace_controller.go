package controllers

import (
	"context"
	"fmt"
	"net/http"

	"kute_k8s_apis/internal/models"

	"github.com/gin-gonic/gin"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetNamespaces responds with the list of all namespaces as JSON.
func GetNamespaces(c *gin.Context) {
	println("GetNamespaces")
	clientset, err := models.NewClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	namespaces, err := models.GetNamespaces(clientset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, namespaces)
}

// CreateNamespace creates a new namespace.
func CreateNamespace(c *gin.Context) {
	clientset, err := models.NewClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var requestBody struct {
		Name string `json:"name"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.CreateNamespace(clientset, requestBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": fmt.Sprintf("Namespace %s created successfully", requestBody.Name), "status": true})
}

func GetNamespace(c *gin.Context) {
	namespaceName := c.Param("name")

	// Check if the namespace exists
	exists, err := NamespaceExists(namespaceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Namespace %s exists", namespaceName), "status": true})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Namespace %s does not exist", namespaceName), "status": false})
	}
}

// NamespaceExists checks if a namespace with the given name exists
func NamespaceExists(namespaceName string) (bool, error) {
	// Assuming you have a Kubernetes client instance named 'clientset'
	clientset, err := models.NewClient()
	if err != nil {
		return false, err
	}
	_, err = clientset.CoreV1().Namespaces().Get(context.TODO(), namespaceName, v1.GetOptions{})
	if err == nil {
		// Namespace exists
		return true, nil
	}

	// Check if the error is due to the namespace not existing
	if IsNotFound(err) {
		// Namespace does not exist
		return false, nil
	}

	// An unexpected error occurred
	return false, err
}

// IsNotFound checks if the error is a Kubernetes NotFound error

func IsNotFound(err error) bool {
	if statusError, isStatus := err.(*errors.StatusError); isStatus {
		return statusError.ErrStatus.Reason == metav1.StatusReasonNotFound
	}
	return false
}

// internal/controllers/your_controller.go

// DeleteNamespace deletes a namespace by name
func DeleteNamespace(c *gin.Context) {
	namespaceName := c.Param("name")
	println("DeleteNamespace", namespaceName)
	// Get the Kubernetes clientset
	clientset, err := models.NewClient()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the namespace exists
	exists, err := NamespaceExists(namespaceName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if exists {
		// Delete the namespace
		err := models.DeleteNamespace(clientset, namespaceName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Namespace %s deleted successfully", namespaceName)})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Namespace %s does not exist", namespaceName)})
	}
}
