package daemonset

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

// Lister is a (k8s.io/kubernetes/pkg/client/unversioned).DaemonSetInterface compatible
// interface designed only for listing daemonSets. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Lister interface {
	List(api.ListOptions) (*extensions.DaemonSetList, error)
}

// FakeLister is a Lister implementation designed to be used in unit tests
type FakeLister struct {
	DaemonSetList *extensions.DaemonSetList
	Err           error
}

// List is the Lister interface implementation. It simply returns f.DaemonSetList, f.Err
func (f FakeLister) List(api.ListOptions) (*extensions.DaemonSetList, error) {
	return f.DaemonSetList, f.Err
}
