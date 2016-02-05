package kubeapp

import (
	"fmt"
	"os"
)

type ErrNoSuchEnvVar struct {
	varName string
}

func (e ErrNoSuchEnvVar) Error() string {
	return fmt.Sprintf("environment variable %s doesn't exist", e.varName)
}

func getEnv(name string) (string, error) {
	n := os.Getenv(name)
	if n == "" {
		return "", ErrNoSuchEnvVar{varName: name}
	}
	return n, nil
}
