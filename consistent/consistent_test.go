package consistent

import (
	"log"
	"os"
	"testing"
)

func TestWHEN_AddNodeWithReplicasCalledForConsistentHashFunction_THEN_MatchNumberOfReplicas(t *testing.T) {
	h := &ConsistentHashing{
		Nodes:    make(map[uint32]string),
		Replicas: 3,
		Keys:     make([]uint32, 0),
		Logger:   log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	h.AddNode("node1")

	if len(h.Nodes) != 3 {
		t.Errorf("Expected 3 nodes, but got %d", len(h.Nodes))
	}
}

func TestWHEN_AddNodeWithReplicasCalledForConsistentHashFunction_THEN_MatchSameEventToSameReplica(t *testing.T) {
	h := &ConsistentHashing{
		Nodes:    make(map[uint32]string),
		Replicas: 3,
		Keys:     make([]uint32, 0),
		Logger:   log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
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
		Nodes:    make(map[uint32]string),
		Replicas: 0,
		Logger:   log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	requestedNode := "hello"

	h.AddNode("Initial")
	h.AddNode("test")
	h.AddNode("TestForConsistentHashing")
	h.AddNode("Essai")

	expected, err1 := h.Hash(requestedNode, 0)

	if err1 != nil {
		t.Errorf("Expected no errors to occur but got %s", err1)
	}

	// remove a bunch of nodes
	h.RemoveNode("Essai")
	h.RemoveNode("test")

	actual, err2 := h.Hash(requestedNode, 0)

	if err2 != nil {
		t.Errorf("Expected no errors to occur but got %s", err2)
	}

	if expected != actual && expected != "" {
		t.Errorf("Expected the Node to be `%s` but actual Node was `%s` for key `%s`", expected, actual, requestedNode)
	}
}

func TestWHEN_providedWithEmptyUUID_THEN_ReturnEmptyString(t *testing.T) {
	h := &ConsistentHashing{
		Nodes:    make(map[uint32]string),
		Replicas: 0,
		Logger:   log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	requestedNode := ""

	h.AddNode("Initial")
	h.AddNode("test")
	h.AddNode("TestForConsistentHashing")
	h.AddNode("Essai")

	expected, err := h.Hash(requestedNode, 0)

	if err == nil || len(expected) != 0 {
		t.Errorf("Expected the returned node to be empty string, but got `%s`", expected)
	}
}

func TestWHEN_SetReplicas_THEN_ReplicasCorrectlySet(t *testing.T) {
	h := &ConsistentHashing{
		Nodes:    make(map[uint32]string),
		Replicas: 0,
		Logger:   log.New(os.Stdout, "hashProfiler: ", log.LstdFlags),
	}

	h.SetReplicas(100)
	h.AddNode("test")

	if len(h.Nodes) != 100 {
		t.Errorf("Expected the number of nodes to be a factor of the number of replicas, but got %d", len(h.Nodes))
	}

	h.SetReplicas(1000)
	h.AddNode("test-server-x")
	h.AddNode("test-server-y")

	if len(h.Nodes) != 2100 {
		t.Errorf("Expected the number of nodes to be a factor of the number of replicas, but got %d", len(h.Nodes))
	}
}
