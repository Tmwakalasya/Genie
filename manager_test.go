package main

import (
	"errors"
	"testing"
)

func TestManagerEnforcesQuota(t *testing.T) {
	// One owner cannot exceed the configured quota, and rejection does not add another VM.
	store := NewStore()
	manager := NewManager(store, 2)
	_, firstError := manager.requestVM("Tuntu")
	_, secondError := manager.requestVM("Tuntu")

	if firstError != nil {
		t.Fatalf("first request returned an unexpected error: %v", firstError)
	}

	if secondError != nil {
		t.Fatalf("second request returned an unexpected error: %v", secondError)
	}
	_, thirdError := manager.requestVM("Tuntu")

	if !errors.Is(thirdError, ErrQuotaExceeded) {
		t.Fatalf("expected ErrQuotaExceeded, got %v", thirdError)
	}

	got := len(store.List("Tuntu"))
	want := 2

	if got != want {
		t.Fatalf("expected %d VMs, got %d", want, got)
	}
}

func TestQuotaIsPerOwner(t *testing.T) {
	// One owner filling their quota does not prevent another owner from creating a VM.
	store := NewStore()
	manager := NewManager(store, 2)
	_, firstErr := manager.requestVM("Tuntu")
	_, secondErr := manager.requestVM("Tuntu")
	_, aliceError := manager.requestVM("Alice")
	if firstErr != nil {
		t.Fatalf("Tuntu's first request returned an unexpected error: %v", firstErr)
	}
	if secondErr != nil {
		t.Fatalf("Tuntu's second request returned an unexpected error: %v", secondErr)
	}
	if aliceError != nil {
		t.Fatalf("Alice's first request should succeed %v", aliceError)
	}
	_, thirdErr := manager.requestVM("Tuntu")
	if !errors.Is(thirdErr, ErrQuotaExceeded) {
		t.Fatalf("expected ErrQuotaExceeded, got %v", thirdErr)
	}
	got := len(store.List("Alice"))
	want := 1

	if got != want {
		t.Fatalf("expected Alice to have %d VM, got %d", want, got)
	}
}

func TestManagerTerminatedVMDoesNotConsumeQuota(t *testing.T) {
	store := NewStore()
	manager := NewManager(store, 1)
	owner := "Tuntu"

	store.Add(VM{
		Owner: owner,
		State: Terminated,
	})

	newVM, err := manager.requestVM(owner)

	if err != nil {
		t.Fatalf(
			"expected request to succeed with only a terminated VM stored, got %v",
			err,
		)
	}

	if newVM.State != Running {
		t.Fatalf("expected new VM state %q, got %q", Running, newVM.State)
	}

	got := len(store.List(owner))
	want := 2

	if got != want {
		t.Fatalf("expected %d stored VM records, got %d", want, got)
	}
}
