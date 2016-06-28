package node

import (
	"k8s.io/kubernetes/pkg/api"
)

// Getter is a (k8s.io/kubernetes/pkg/client/unversioned).NodeInterface compatible
// interface designed only for getting nodes. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Getter interface {
	Get(name string) (*api.Node, error)
}

// FakeGetter is a Getter implementation designed to be used in unit tests
type FakeGetter struct {
	Node *api.Node
	Err  error
}

// Get is the Getter interface implementation. It simply returns f.RCList, f.Err
func (f FakeGetter) Get(name string) (*api.Node, error) {
	return f.Node, f.Err
}
