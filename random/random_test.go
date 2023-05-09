package random

import (
	"log"
	"os"
	"testing"
)

type Tuple struct {
	first  string
	second int
}

func TestWHEN_HashFunctionCalledWithNullEvent_THEN_NullPointerExceptionThrown(t *testing.T) {
	hasher := &RandomHashing{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	event := ""
	n := 3

	_, err := hasher.Hash(event, n)
	if err == nil {
		t.Error("Expected non-nil error as event is empty but got nil")
	}
}

func TestWHEN_HashFunctionCalledWithNullShards_THEN_NullPointerExceptionThrown(t *testing.T) {
	hasher := &RandomHashing{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	event := "1"
	n := 0

	_, err := hasher.Hash(event, n)
	if err == nil {
		t.Error("Expected non-nil error as shards number is 0 but got nil")
	}
}

func TestWHEN_HashFunctionCalledWithKeyAndShardNumbers_THEN_ResultMatchesExpected(t *testing.T) {
	hasher := &RandomHashing{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	eventList := []Tuple{
		{"1Test", 1},
		{"2Hello", 4},
		{"Test 3", 5},
		{"hello", 2},
	}

	for _, v := range eventList {
		result, err := hasher.Hash(v.first, v.second)

		if err == nil {
			if result == "" {
				t.Errorf("Hash(%s, %d) = %s; expected != \"\"", v.first, v.second, result)
			}
		}
	}
}

func TestRandomHashing_AddNode(t *testing.T) {
	h := RandomHashing{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}
	expectedError := "AddNode method is not implemented for RandomHashing"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("AddNode did not panic with error message '%s'", expectedError)
		} else if r != expectedError {
			t.Errorf("AddNode panicked with error message '%s', but expected '%s'", r, expectedError)
		}
	}()

	h.AddNode("node")
}

func TestRandomHashing_RemoveNode(t *testing.T) {
	h := RandomHashing{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	expectedError := "RemoveNode method is not implemented for RandomHashing"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("RemoveNode did not panic with error message '%s'", expectedError)
		} else if r != expectedError {
			t.Errorf("RemoveNode panicked with error message '%s', but expected '%s'", r, expectedError)
		}
	}()

	h.RemoveNode("node")
}

func TestRandomHashing_SetReplicas(t *testing.T) {
	h := RandomHashing{
		Logger: log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	expectedError := "SetReplicas method is not implemented for RandomHashing"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("SetReplicas did not panic with error message '%s'", expectedError)
		} else if r != expectedError {
			t.Errorf("SetReplicas panicked with error message '%s', but expected '%s'", r, expectedError)
		}
	}()

	h.SetReplicas(4)
}
