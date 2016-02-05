package kubeapp

import (
	kubeclient "github.com/jchauncey/kubeclient/client/unversioned"
)

// ClientInCluster is a convenience function for calling NewInCluster inside the github.com/jchauncey/kubeclient/client/unversioned package
func ClientInCluster() (kubeclient.Interface, error) {
	return kubeclient.NewInCluster()
}
