package node

import (
	"k8s.io/kubernetes/pkg/api"
)

// Lister is a (k8s.io/kubernetes/pkg/client/unversioned).NodeInterface compatible
// interface designed only for listing nodes. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Lister interface {
	List(api.ListOptions) (*api.NodeList, error)
}

// FakeLister is a Lister implementation designed to be used in unit tests
type FakeLister struct {
	NodeList *api.NodeList
	Err      error
}

// List is the Lister interface implementation. It simply returns f.RCList, f.Err
func (f FakeLister) List(api.ListOptions) (*api.NodeList, error) {
	return f.NodeList, f.Err
}
