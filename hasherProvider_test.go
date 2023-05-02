package hasherProvider

import (
	"testing"
)

func TestWHEN_requestedAlgoDoesNotExist_THEN_returnError(t *testing.T) {
	hp := HasherProvider{}

	algo := 3
	_, err := hp.GetHasher(algo)
	if err == nil {
		t.Errorf("Expected error for invalid algorithm type %d, but got nil", algo)
	}
}

func TestWHEN_requestForConsistentHasher_THEN_NoError(t *testing.T) {
	hp := HasherProvider{}

	algo := CONSISTENT_HASHING
	hasher, err := hp.GetHasher(algo)
	if hasher != nil && err != nil {
		t.Errorf("Unexpected error for valid algorithm type %d: %v", algo, err)
	}
}

func TestWHEN_requestForRandomHasher_THEN_noError(t *testing.T) {
	hp := HasherProvider{}

	algo := RANDOM_HASHING
	hasher, err := hp.GetHasher(algo)
	if hasher != nil && err != nil {
		t.Errorf("Unexpected error for valid algorithm type %d: %v", algo, err)
	}
}

func TestWHEN_requestForUniformHasher_THEN_NoError(t *testing.T) {
	hp := HasherProvider{}

	algo := UNIFORM_HASHING
	hasher, err := hp.GetHasher(algo)
	if hasher != nil && err != nil {
		t.Errorf("Unexpected error for valid algorithm type %d: %v", algo, err)
	}
}
