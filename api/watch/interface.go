package util

import (
	"errors"
	"sync/atomic"

	"k8s.io/kubernetes/pkg/watch"
)

var (
	ErrStopped = errors.New("stopped")
)

// FakeInterface is a fake watch.Interface implementation to be used in unit tests. FakeInterface implementations have a Produce func defined on them, which allows callers to send a single watch.Event to one receiver on a channel returned by ResultChan(). Note that this implements a 1:1 send - if there are multiple receivers, only one will get the produced event
type FakeInterface struct {
	eventsCh chan watch.Event
	closed   uint32
	stopCh   chan struct{}
}

// NewFakeInterface returns a new FakeInterface instance
func NewFakeInterface() *FakeInterface {
	return &FakeInterface{
		eventsCh: make(chan watch.Event),
		closed:   uint32(1),
		stopCh:   make(chan struct{}),
	}
}

// ResultChan is the watch.Interface implementation
func (i *FakeInterface) ResultChan() <-chan watch.Event {
	return i.eventsCh
}

// Stop is the watch.Interface implementation
func (i *FakeInterface) Stop() {
	close(i.eventsCh)
	atomic.SwapUint32(&i.closed, 1)
}

// Produce produces an event that will be received on channels returned by ResultChan. Returns ErrStopped if Stop has been called, and blocks until a receive has happened on exactly one ResultChan channel
func (i *FakeInterface) Produce(evt watch.Event) error {
	if atomic.LoadUint32(&i.closed) == 1 {
		return ErrStopped
	}
	i.eventsCh <- evt
	return nil
}
