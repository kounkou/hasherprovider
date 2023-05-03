package uniform

import (
	"testing"
	"fmt"
)

type Tuple struct {
	first  string
	second int
	third  string
}

func TestWHEN_HashFunctionCalledWithNullEvent_THEN_NullPointerExceptionThrown(t *testing.T) {
	hasher := &UniformHashing{}

	event := ""
	n := 3

	_, err := hasher.Hash(event, n)
	if err == nil {
		t.Error("Expected non-nil error as event is empty but got nil")
	}
}

func TestWHEN_HashFunctionCalledWithNullShards_THEN_NullPointerExceptionThrown(t *testing.T) {
	hasher := &UniformHashing{}

	event := "1"
	n := 0

	_, err := hasher.Hash(event, n)
	if err == nil {
		t.Error("Expected non-nil error as shards number is 0 but got nil")
	}
}

func TestWHEN_HashFunctionCalledWithKeyAndShardNumbers_THEN_ResultMatchesExpected(t *testing.T) {
	hasher := &UniformHashing{}

	eventList := []Tuple{
		{"1Test", 1, "0"},
		{"2Hello", 4, "2"},
		{"Test 3", 5, "4"},
		{"hello", 2, "0"},
	}

	for _, v := range eventList {
		result, err := hasher.Hash(v.first, v.second)

		if err == nil {
			if result != v.third {
				t.Errorf("Hash(%s, %d) = %s; expected %s", v.first, v.second, result, v.third)
			}
		}
	}
}

func TestUniformHashing_AddNode(t *testing.T) {
    h := UniformHashing{}
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

func TestUniformHashing_RemoveNode(t *testing.T) {
    h := UniformHashing{}
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

func TestUniformHashing_SetReplicas(t *testing.T) {
    REPLICAS_EXPECTED := 3

    h := UniformHashing{}
    h.SetReplicas(REPLICAS_EXPECTED)

    if h.Replicas != REPLICAS_EXPECTED {
        t.Error("Expected replicas to be", REPLICAS_EXPECTED)
    }
}
