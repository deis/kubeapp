package namespace

import (
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/watch"
)

// Watcher is a (k8s.io/kubernetes/pkg/client/unversioned).NamespaceInterface compatible
// interface which only has the Watch function. It's used in places that only need perform watches,
// to make those codebases easier to test and more easily swappable with other implementations
// (should the need arise).
//
// Example usage:
//
//	var nsl NamespaceWatcher
//	nsl = kubeClient.Namespaces()
type Watcher interface {
	Watch(label labels.Selector, field fields.Selector, resourceVersion string) (watch.Interface, error)
}
