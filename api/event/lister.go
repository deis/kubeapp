package event

import "k8s.io/kubernetes/pkg/api"

// Lister is a (k8s.io/kubernetes/pkg/client/unversioned).EventInterface compatible
// interface designed only for listing events. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Lister interface {
	List(api.ListOptions) (*api.EventList, error)
}

// FakeLister is a Lister implementation designed to be used in unit tests
type FakeLister struct {
	EventList *api.EventList
	Err       error
}

// List is the Lister interface implementation. It simply returns f.EventList, f.Err
func (f FakeLister) List(api.ListOptions) (*api.EventList, error) {
	return f.EventList, f.Err
}
