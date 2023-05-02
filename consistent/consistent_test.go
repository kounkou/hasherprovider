package consistent

import (
	"testing"
)

func TestWHEN_AddNodeWithReplicasCalledForConsistentHashFunction_THEN_MatchNumberOfReplicas(t *testing.T) {
	h := &ConsistentHashing{
		replicas:    3,
		nodes:       make(map[uint32]string),
		sortedNodes: make([]uint32, 0),
	}

	h.AddNode("node1")

	if len(h.nodes) != 3 {
		t.Errorf("Expected 3 nodes, but got %d", len(h.nodes))
	}

	if len(h.sortedNodes) != 3 {
		t.Errorf("Expected 3 sorted nodes, but got %d", len(h.sortedNodes))
	}
}

func TestWHEN_AddNodeWithReplicasCalledForConsistentHashFunction_THEN_MatchSameEventToSameReplica(t *testing.T) {
	h := &ConsistentHashing{
		replicas:    3,
		nodes:       make(map[uint32]string),
		sortedNodes: make([]uint32, 0),
	}

	h.AddNode("node1")
	result1 := h.GetImmediateNode("node1")
	h.AddNode("node2")
	result2 := h.GetImmediateNode("node1")

	if result1 != result2 {
		t.Errorf("Expected the Node to be the same for the same key after adding new different node")
	}
}
