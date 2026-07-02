package main

import "sync"

// Manager coordinates the lifecycle of virtual machines within the control plane.
// It enforces resource quotas and orchestrates requests to the underlying Store.
type Manager struct {
	mu         sync.Mutex
	store      *Store
	quotaLimit int // The maximum number of active VMs allowed per owner.
}

// NewManager initializes a new control plane brain with a specified quota limit.
func NewManager(store *Store, quotaLimit int) *Manager {
	return &Manager{
		store:      store,
		quotaLimit: quotaLimit,
	}
}

// isAlive determines if a VM is currently operational.
// A VM is considered alive if it hasn't entered a terminal state ("failed" or "terminated").
func (v VM) isAlive() bool {
	return v.State != "failed" && v.State != "terminated"
}

// aliveCount calculates the total number of non-terminal VMs owned by a user.
// This count is strictly used to validate and enforce user quota allocations.
func (m *Manager) aliveCount(owner string) int {
	m.mu.Lock()
	defer m.mu.Unlock()

	count := 0
	for _, vm := range m.store.List(owner) {
		if vm.isAlive() {
			count++
		}
	}
	return count
}
