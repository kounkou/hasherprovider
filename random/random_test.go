package random

import (
	"testing"
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

    eventList := []Tuple {
        Tuple{ "1Test",  1 },
        Tuple{ "2Hello", 4 },
        Tuple{ "Test 3", 5 },
        Tuple{ "hello",  2 },
    }

    for _, v := range eventList {
        result, err := hasher.Hash(v.first, v.second)

        if err == nil {
            if result >= v.second {
                t.Errorf("Hash(%s, %d) = %d; expected <= %d", v.first, v.second, result, v.second)
            }
        }
    }
}
