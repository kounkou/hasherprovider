package random

import (
	"testing"
	"fmt"
)

type Tuple struct {
	first  string
	second int
}

func TestWHEN_HashFunctionCalledWithNullEvent_THEN_NullPointerExceptionThrown(t *testing.T) {
	hasher := &RandomHashing{}

	event := ""
	n := 3

	_, err := hasher.Hash(event, n)
	if err == nil {
		t.Error("Expected non-nil error as event is empty but got nil")
	}
}

func TestWHEN_HashFunctionCalledWithNullShards_THEN_NullPointerExceptionThrown(t *testing.T) {
	hasher := &RandomHashing{}

	event := "1"
	n := 0

	_, err := hasher.Hash(event, n)
	if err == nil {
		t.Error("Expected non-nil error as shards number is 0 but got nil")
	}
}

func TestWHEN_HashFunctionCalledWithKeyAndShardNumbers_THEN_ResultMatchesExpected(t *testing.T) {
	hasher := &RandomHashing{}

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
    h := RandomHashing{}
    err := func() (err error) {
        defer func() {
            if r := recover(); r != nil {
                err = fmt.Errorf("AddNode panicked: %v", r)
            }
        }()
        h.AddNode("some-string-value")
        return
    }()
    if err != nil {
        t.Error(err)
    }
}

func TestRandomHashing_RemoveNode(t *testing.T) {
    h := RandomHashing{}
    err := func() (err error) {
        defer func() {
            if r := recover(); r != nil {
                err = fmt.Errorf("AddNode panicked: %v", r)
            }
        }()
        h.RemoveNode("some-string-value")
        return
    }()
    if err != nil {
        t.Error(err)
    }
}

func TestRandomHashing_SetReplicas(t *testing.T) {
    REPLICAS_EXPECTED := 3

    h := RandomHashing{}
    h.SetReplicas(REPLICAS_EXPECTED)

    if h.Replicas != REPLICAS_EXPECTED {
        t.Error("Expected replicas to be", REPLICAS_EXPECTED)
    }
}