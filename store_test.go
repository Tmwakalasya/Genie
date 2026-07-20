package main

import (
	"errors"
	"testing"
)

func TestStoreUpdateChangesState(t *testing.T) {
	store := NewStore()
	created := store.Add(VM{
		Owner: "Tuntu",
		State: Running,
	})

	err := store.Update(created.ID, Terminated)
	if err != nil {
		t.Fatalf("Update returned unexpected error: %v", err)
	}

	got, ok := store.Get(created.ID)
	if !ok {
		t.Fatalf("expected VM %d to exist after Update", created.ID)
	}
	if got.State != Terminated {
		t.Fatalf("expected state %q, got %q", Terminated, got.State)
	}
}

func TestStoreUpdateMissingID(t *testing.T) {
	store := NewStore()

	err := store.Update(999, Terminated)
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
