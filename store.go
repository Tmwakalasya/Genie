package main

import (
	"errors"
	"sync"
)

var ErrNotFound = errors.New("Resource not found")

// Store is the in-memory, thread-safe home of all VM state.
type Store struct {
	mu     sync.Mutex
	nextID uint
	vms    map[uint]VM
}

func NewStore() *Store {
	return &Store{
		vms: make(map[uint]VM),
	}
}

// Add mints an ID for the VM, files it, and returns the completed VM.
func (s *Store) Add(vm VM) VM {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.nextID++
	vm.ID = s.nextID
	s.vms[vm.ID] = vm

	return vm
}

// Get retrieves a VM by ID. The bool reports whether it was found.
func (s *Store) Get(id uint) (VM, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	vm, ok := s.vms[id]
	return vm, ok
}

// List returns copies of all VMs belonging to owner.
func (s *Store) List(owner string) []VM {
	s.mu.Lock()
	defer s.mu.Unlock()

	var results []VM
	for _, vm := range s.vms {
		if vm.Owner == owner {
			results = append(results, vm)
		}
	}
	return results
}

func (s *Store) Update(id uint, state State) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	vm, ok := s.vms[id]
	if !ok {
		return ErrNotFound
	}
	vm.State = state
	s.vms[id] = vm
	return nil

}
