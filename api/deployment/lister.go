package deployment

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

// Lister is a (k8s.io/kubernetes/pkg/client/unversioned).DeploymentInterface compatible
// interface designed only for listing deployments. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Lister interface {
	List(api.ListOptions) (*extensions.DeploymentList, error)
}

// FakeLister is a Lister implementation designed to be used in unit tests
type FakeLister struct {
	DeploymentList *extensions.DeploymentList
	Err            error
}

// List is the Lister interface implementation. It simply returns f.DeploymentList, f.Err
func (f FakeLister) List(api.ListOptions) (*extensions.DeploymentList, error) {
	return f.DeploymentList, f.Err
}
