package secret

import (
	"k8s.io/kubernetes/pkg/api"
)

// Lister is a (k8s.io/kubernetes/pkg/client/unversioned).SecretsInterface compatible
// interface designed only for listing secrets. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Lister interface {
	List(api.ListOptions) (*api.SecretList, error)
}

// FakeLister is a Lister implementation to be used in unit tests
type FakeLister struct {
	Secrets *api.SecretList
	Err     error
}

// List is the Lister interface implementation. It just returns f.Secrets, f.Err
func (f FakeLister) List(api.ListOptions) (*api.SecretList, error) {
	return f.Secrets, f.Err
}
