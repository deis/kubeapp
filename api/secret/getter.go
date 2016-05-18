package secret

import (
	"k8s.io/kubernetes/pkg/api"
)

// Getter is a (k8s.io/kubernetes/pkg/client/unversioned).SecretsInterface compatible
// interface designed only for getting a secret. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Getter interface {
	Get(string) (*api.Secret, error)
}

// FakeGetter is a Getter implementation to be used in unit tests
type FakeGetter struct {
	Secret *api.Secret
	Err    error
}

// Get is the Getter interface implementation. It just returns f.Svc, f.Err
func (f FakeGetter) Get(name string) (*api.Secret, error) {
	return f.Secret, f.Err
}
