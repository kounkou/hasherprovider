package hasherprovider

import (
	"log"
	"os"
	"testing"
)

func TestWHEN_requestedAlgoDoesNotExist_THEN_returnError(t *testing.T) {
	hp := HasherProvider{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	algo := 3
	_, err := hp.GetHasher(algo)
	if err == nil {
		t.Errorf("Expected error for invalid algorithm type %d, but got nil", algo)
	}
}

func TestWHEN_requestForConsistentHasher_THEN_NoError(t *testing.T) {
	hp := HasherProvider{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	algo := CONSISTENT_HASHING
	hasher, err := hp.GetHasher(algo)
	if hasher != nil && err != nil {
		t.Errorf("Unexpected error for valid algorithm type %d: %v", algo, err)
	}
}

func TestWHEN_requestForRandomHasher_THEN_noError(t *testing.T) {
	hp := HasherProvider{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	algo := RANDOM_HASHING
	hasher, err := hp.GetHasher(algo)
	if hasher != nil && err != nil {
		t.Errorf("Unexpected error for valid algorithm type %d: %v", algo, err)
	}
}

func TestWHEN_requestForUniformHasher_THEN_NoError(t *testing.T) {
	hp := HasherProvider{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	algo := UNIFORM_HASHING
	hasher, err := hp.GetHasher(algo)
	if hasher != nil && err != nil {
		t.Errorf("Unexpected error for valid algorithm type %d: %v", algo, err)
	}
}

func TestWHEN_fullFlow_THEN_Success(t *testing.T) {
	hp := HasherProvider{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	algo := CONSISTENT_HASHING
	hasher, err := hp.GetHasher(algo)

	if err != nil {
		t.Errorf("Unexpected error for valid algorithm type %d : %v", algo, err)
	}

	hasher.AddNode("1")
	hasher.AddNode("5")
	hasher.AddNode("8")
	hasher.RemoveNode("8")

	_, err = hasher.Hash("test", 4)

	if err != nil {
		t.Errorf("Unexpected error for valid Hashing %d : %v", algo, err)
	}
}

func TestWHEN_logIsNotSet_THEN_NoPanic(t *testing.T) {
	hp := HasherProvider{
		Logger: nil,
	}

	algo := CONSISTENT_HASHING

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Function panicked with %v", r)
		}
	}()

	hasher, _ := hp.GetHasher(algo)

	_, err := hasher.Hash("test", 4)

	hasher.AddNode("1")
	hasher.AddNode("5")
	hasher.AddNode("8")
	hasher.RemoveNode("8")

	if err != nil {
		t.Errorf("Unexpected error for valid Hashing %d : %v", algo, err)
	}
}
