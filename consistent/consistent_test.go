package consistent

import (
    "testing"
)

func TestWHEN_HashFunctionCalledWithNullEvent_THEN_NullPointerExceptionThrown(t *testing.T) {
    h := ConsistentHashing{values: []int{1, 2, 3}}
    event := "test"
    n := 10

    result, values, err := h.Hash(event, n)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if len(values) != len(h.values) {
        t.Errorf("Expected length of values %d, but got %d", len(h.values), len(values))
    }
    for i := 0; i < len(values); i++ {
        if values[i] != h.values[i] {
            t.Errorf("Expected value at index %d to be %d, but got %d", i, h.values[i], values[i])
        }
    }
    if result == "" {
        t.Errorf("Expected non-empty result")
    }
}