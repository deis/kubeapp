package secret

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/watch"
)

// Watcher is a (k8s.io/kubernetes/pkg/client/unversioned).SecretsInterface compatible
// interface designed only for watching secrets. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Watcher interface {
	Watch(api.ListOptions) (watch.Interface, error)
}

// FakeWatcher is a Watcher implementation to be used in unit tests
type FakeWatcher struct {
	IFace watch.Interface
	Err   error
}

// Watch is the Watcher interface implementation
func (f FakeWatcher) Watch(api.ListOptions) (watch.Interface, error) {
	return f.IFace, f.Err
}
