package secret

import (
	"k8s.io/kubernetes/pkg/api"
)

// Creator is a (k8s.io/kubernetes/pkg/client/unversioned).SecretsInterface compatible
// interface designed only for creating a secret. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Creator interface {
	Create(*api.Secret) (*api.Secret, error)
}

// FakeCreator is a Creator implementation to be used in unit tests
type FakeCreator struct {
	Created    []*api.Secret
	Errs       []error
	CreateFunc func(*api.Secret) (*api.Secret, error)
}

// Create is the Creator interface implementation. It just returns f.Svc, f.Err
func (f *FakeCreator) Create(sec *api.Secret) (*api.Secret, error) {
	ret, err := f.CreateFunc(sec)
	if err != nil {
		f.Errs = append(f.Errs, err)
		return nil, err
	}
	f.Created = append(f.Created, ret)
	return ret, nil
}
