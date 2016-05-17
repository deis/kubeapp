package namespace

// Deleter is a (k8s.io/kubernetes/pkg/client/unversioned).NamespaceInterface compatible
// interface designed only for deleting namespaces. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Deleter interface {
	Delete(name string) error
}

// FakeDeleter is a NamespaceDeleter implementation to be used in unit tests
type FakeDeleter struct {
	NsDeleted []string
	Err       error
}

// Delete is the Deleter interface implementation. It just returns f.Err
func (f *FakeDeleter) Delete(name string) error {
	f.NsDeleted = append(f.NsDeleted, name)
	return f.Err
}
