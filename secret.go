package kubeapp

import (
	"fmt"
	"io/ioutil"
)

type Secret interface {
	fmt.Stringer
	Value() string
}

type staticSecret struct {
	val string
}

func (f *staticSecret) Value() string  { return f.val }
func (f *staticSecret) String() string { return f.val }

func SecretFromFile(filename string) (Secret, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &staticSecret{val: string(b)}, nil
}

func StaticSecret(val string) Secret {
	return &staticSecret{val: val}
}
