package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var ErrNotFound = errors.New("no service addresses found")

// service discovery logic should be technology agnostic,
// so our services are not tied and hardly coupled to it.
type Registry interface {
	Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error
	Deregister(ctx context.Context, instanceID string, serviceName string) error
	ServiceAddress(ctx context.Context, serviceID string) ([]string, error)
	ReportHealthyState(instanceID string, serviceName string) error
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
