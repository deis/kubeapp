package rc

import (
	"k8s.io/kubernetes/pkg/api"
)

// Getter is a (k8s.io/kubernetes/pkg/client/unversioned).ReplicationControllerInterface compatible
// interface designed only for getting replication controllers. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Getter interface {
	Get(name string) (*api.ReplicationController, error)
}

// FakeGetter is a Getter implementation designed to be used in unit tests
type FakeGetter struct {
	RC  *api.ReplicationController
	Err error
}

// Get is the Getter interface implementation. It simply returns f.RCList, f.Err
func (f FakeGetter) Get(api.ListOptions) (*api.ReplicationController, error) {
	return f.RC, f.Err
}
