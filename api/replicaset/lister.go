package replicaset

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
)

// Lister is a (k8s.io/kubernetes/pkg/client/unversioned).ReplicaSetInterface compatible
// interface designed only for listing replicaSets. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Lister interface {
	List(api.ListOptions) (*extensions.ReplicaSetList, error)
}

// FakeLister is a Lister implementation designed to be used in unit tests
type FakeLister struct {
	ReplicaSetList *extensions.ReplicaSetList
	Err            error
}

// List is the Lister interface implementation. It simply returns f.ReplicaSetList, f.Err
func (f FakeLister) List(api.ListOptions) (*extensions.ReplicaSetList, error) {
	return f.ReplicaSetList, f.Err
}
