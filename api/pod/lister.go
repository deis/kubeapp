package pod

import (
	"k8s.io/kubernetes/pkg/api"
)

// Lister is a (k8s.io/kubernetes/pkg/client/unversioned).PodInterface compatible
// interface designed only for listing pods. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Lister interface {
	List(api.ListOptions) (*api.PodList, error)
}

// FakeLister is a Lister implementation designed to be used in unit tests
type FakeLister struct {
	PodList *api.PodList
	Err     error
}

// List is the Lister interface implementation. It simply returns f.PodList, f.Err
func (f FakeLister) List(api.ListOptions) (*api.PodList, error) {
	return f.PodList, f.Err
}
