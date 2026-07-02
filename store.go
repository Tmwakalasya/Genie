package main

import (
	"sync"
)

// Store provides an in-memory, thread-safe repository for managing VM states.
type Store struct {
	mu  sync.Mutex
	vms map[uint]*VM // Storing pointers prevents copying the internal VM mutex.
}

// NewStore initializes a Store with an empty memory allocation map.
func NewStore() *Store {
	return &Store{
		vms: make(map[uint]*VM),
	}
}

// AddVM registers a new VM instance or overwrites an existing entry.
func (s *Store) AddVM(vm *VM) {
	if vm == nil {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.vms[vm.ID] = vm
}

// Get retrieves a specific VM instance by its unique identifier.
func (s *Store) Get(id uint) (*VM, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	vm, ok := s.vms[id]
	return vm, ok
}

// List collects all active and historical VMs associated with a single owner.
// If no VMs are found, it returns nil to explicitly signal an empty allocation state.
func (s *Store) List(owner string) []*VM {
	s.mu.Lock()
	defer s.mu.Unlock()

	var results []*VM
	for _, vm := range s.vms {
		if vm.Owner == owner {
			results = append(results, vm)
		}
	}

	// Retaining the slice allocation logic. Returning nil explicitly
	// distinguishes "no records found" from an initialized container.
	return results
}
