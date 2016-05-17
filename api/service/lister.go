package service

import (
	"k8s.io/kubernetes/pkg/api"
)

// Lister is a (k8s.io/kubernetes/pkg/client/unversioned).ServiceInterface compatible
// interface designed only for listing services. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Lister interface {
	List(opts api.ListOptions) (*api.ServiceList, error)
}

// FakeLister is a ServiceLister implementation to be used in unit tests
type FakeLister struct {
	SvcList *api.ServiceList
	Err     error
}

// List is the Lister interface implementation. It just returns f.SvcList, f.Err
func (f FakeLister) List(opts api.ListOptions) (*api.ServiceList, error) {
	return f.SvcList, f.Err
}
