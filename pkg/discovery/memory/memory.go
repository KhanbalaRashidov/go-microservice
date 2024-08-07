package memory

import (
	"context"
	"errors"
	"go-microservice/pkg/discovery"
	"sync"
	"time"
)

type Registry struct {
	sync.RWMutex
	serviceAddrs map[string]map[string]*serviceInstance
}

type serviceInstance struct {
	hostPort   string
	lastActive time.Time
}

func NewRegistry() *Registry {
	return &Registry{serviceAddrs: map[string]map[string]*serviceInstance{}}
}

func (r *Registry) Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[serviceName]; !ok {
		r.serviceAddrs[serviceName] =
			map[string]*serviceInstance{}
	}

	r.serviceAddrs[serviceName][instanceID] =
		&serviceInstance{hostPort: hostPort, lastActive: time.
			Now()}
	return nil
}

func (r *Registry) Deregister(ctx context.Context, instanceID string, serviceName string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return nil
	}
	delete(r.serviceAddrs[serviceName], instanceID)
	return nil
}

func (r *Registry) ReportHealthyState(instanceID string, serviceName string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.serviceAddrs[serviceName]; !ok {
		return errors.New("service is not registered yet")
	}
	if _, ok := r.serviceAddrs[serviceName][instanceID]; !ok {
		return errors.New("service instance is not registered yet")
	}

	r.serviceAddrs[serviceName][instanceID].lastActive = time.Now()

	return nil
}

func (r *Registry) ServiceAddresses(ctx context.Context, serviceName string) ([]string, error) {
	r.RLock()
	defer r.RUnlock()
	if len(r.serviceAddrs[serviceName]) == 0 {
		return nil, discovery.ErrNotFound
	}
	var res []string
	for _, i := range r.serviceAddrs[serviceName] {
		if i.lastActive.Before(time.Now().Add(-5 * time.
			Second)) {
			continue
		}
		res = append(res, i.hostPort)
	}
	return res, nil
}
