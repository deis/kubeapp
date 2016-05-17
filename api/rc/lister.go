package rc

import (
	"k8s.io/kubernetes/pkg/api"
)

// Lister is a (k8s.io/kubernetes/pkg/client/unversioned).ReplicationControllerInterface compatible
// interface designed only for listing replication controllers. It should be used as a parameter to functions
// so that they can be more easily unit tested
type Lister interface {
	List(api.ListOptions) (*api.ReplicationControllerList, error)
}
