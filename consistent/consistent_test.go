package consistent

import (
	"testing"
)

func TestWHEN_AddNodeWithReplicasCalledForConsistentHashFunction_THEN_MatchNumberOfReplicas(t *testing.T) {
	h := &ConsistentHashing{
        Nodes:       make(map[uint32]string),
        Replicas:    3,
        Keys:        make([]uint32, 0),
	}

	h.AddNode("node1")

	if len(h.Nodes) != 3 {
		t.Errorf("Expected 3 nodes, but got %d", len(h.Nodes))
	}
}

func TestWHEN_AddNodeWithReplicasCalledForConsistentHashFunction_THEN_MatchSameEventToSameReplica(t *testing.T) {
	h := &ConsistentHashing{
        Nodes:       make(map[uint32]string),
        Replicas:    3,
        Keys:        make([]uint32, 0),
	}

	h.AddNode("node1")
	result1 := h.GetImmediateNode("node1")
	h.AddNode("node2")
	result2 := h.GetImmediateNode("node1")

	if result1 != result2 {
		t.Errorf("Expected the Node to be the same for the same key after adding new different node")
	}
}

func TestWHEN_AddAndRemoveDifferentNodeWithReplicasCalledForConsistentHashFunction_THEN_MatchSameEventToSameReplica(t *testing.T) {
	h := &ConsistentHashing{
        Nodes:       make(map[uint32]string),
        Replicas:    0,
	}

	requestedNode := "hello"

	h.AddNode("Initial")
	h.AddNode("test")
	h.AddNode("TestForConsistentHashing")
	h.AddNode("Essai")

	expected, err1 := h.Hash(requestedNode, 4)

	if err1 != nil {
		t.Errorf("Expected no errors to occur but got %s", err1)
	}

	// remove a bunch of nodes
	h.RemoveNode("Essai")
	h.RemoveNode("test")

	actual, err2 := h.Hash(requestedNode, 4)

	if err2 != nil {
		t.Errorf("Expected no errors to occurr but got %s", err2)
	}

	if expected != actual && expected != "" {
		t.Errorf("Expected the Node to be `%s` but actual Node was `%s` for key `%s`", expected, actual, requestedNode)
	}
}
