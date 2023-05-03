package uniform

import (
	"testing"
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
