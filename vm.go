package main

import (
	"sync"
	"time"
)

// State represents the current operational phase of a virtual machine.
type State string

const (
	Pending    State = "pending"    // VM is scheduled but not yet initialized.
	Booting    State = "booting"    // VM underlying hardware/hypervisor is starting up.
	Running    State = "running"    // VM is active and accepting workloads.
	Failed     State = "failed"     // VM encountered an unrecoverable error.
	Terminated State = "terminated" // VM has been stopped and decommissioned.
)

// VM defines the identity, state, and tracking metrics of a managed virtual machine.
//
// NOTE: This struct contains a sync.Mutex and must always be handled
// via a pointer (*VM) to prevent copying the synchronization state.
type VM struct {
	mu        sync.Mutex // Guards access to mutable fields like State.
	ID        uint
	Owner     string
	State     State
	CreatedAt time.Time
}
