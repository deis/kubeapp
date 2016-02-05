package kubeapp

import (
	"fmt"
	"strconv"
)

type ErrInvalidPort struct {
	portStr string
}

func (e ErrInvalidPort) Error() string {
	return fmt.Sprintf("port %s is invalid", e.portStr)
}

type Service interface {
	Host() string
	Port() int
	HostStr() string
}

// ServiceFromEnvVars looks up the service discovery environment variables described in http://kubernetes.io/v1.1/docs/user-guide/services.html#environment-variables.
//
// Returns an error if any of the environment variables don't exist or any were invalid.
func ServiceFromEnvVars(name string) (Service, error) {
	hst, err := getEnv(fmt.Sprintf("%s_SERVICE_HOST", name))
	if err != nil {
		return nil, err
	}
	prtStr, err := getEnv(fmt.Sprintf("%s_SERVICE_PORT", name))
	if err != nil {
		return nil, err
	}
	prt, err := strconv.Atoi(prtStr)
	if err != nil {
		return nil, ErrInvalidPort{portStr: prtStr}
	}
	return &svcVals{host: hst, port: prt}, nil
}

// StaticService creates a service that contains the given values. It's useful for unit testing in environments outside Kubernetes
func StaticService(host string, port int) Service {
	return &svcVals{host: host, port: port}
}

type svcVals struct {
	host string
	port int
}

func (s *svcVals) Host() string    { return s.host }
func (s *svcVals) Port() int       { return s.port }
func (s *svcVals) HostStr() string { return fmt.Sprintf("%s:%d", s.host, s.port) }
