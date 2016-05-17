package service

import (
	"k8s.io/kubernetes/pkg/api"
)

// Getter is a (k8s.io/kubernetes/pkg/client/unversioned).ServiceInterface compatible
// interface designed only for getting a service. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Getter interface {
	Get(name string) (*api.Service, error)
}

// FakeGetter is a ServiceGetter implementation to be used in unit tests
type FakeGetter struct {
	Svc *api.Service
	Err error
}

// Get is the Getter interface implementation. It just returns f.Svc, f.Err
func (f FakeGetter) Get(name string) (*api.Service, error) {
	return f.Svc, f.Err
}
