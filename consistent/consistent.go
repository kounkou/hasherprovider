// MIT License
//
// Copyright (c) 2023 Godfrain Jacques Kounkou
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package consistent

import (
	"errors"
	"hash/fnv"
	"sort"
	"strconv"
	"log"
)

// With consistent Hashing, the keys already assigned to a shard
// do NOT need to be reassigned. Hence solving the issue introduced
// by the usage of Modulo to be able to perform a consistent Hashing.

type ConsistentHashing struct {
	Nodes    map[uint32]string
	Replicas int
	Keys     []uint32
	Logger   *log.Logger
}

// SetReplicas set the replicas for the entities to be hashed in the ring
func (h *ConsistentHashing) SetReplicas(replicas int) {
	h.Replicas = replicas
}

// AddNode will add a node or entity in the ring using its hashed value
// The ring is then ordered by the hashed value and saved in Keys
func (h *ConsistentHashing) AddNode(node string) {
    h.Logger.Println("[INFO] AddNode ", node)

	for i := 0; i < h.Replicas; i++ {
		key := h.computeHash(node + strconv.Itoa(i))
		h.Nodes[key] = node
		h.Keys = append(h.Keys, key)
	}

	sort.Slice(h.Keys, func(i, j int) bool {
	    return h.Keys[i] < h.Keys[j]
	})
}

// RemoveNode will remove a node or entity from the ring. Then we will also 
// make sure that the Keys are consistent are removal of a node
func (h *ConsistentHashing) RemoveNode(node string) {
    h.Logger.Println("[INFO] RemoveNode ", node)

	for i := 0; i < h.Replicas; i++ {
		key := h.computeHash(node + strconv.Itoa(i))
		delete(h.Nodes, key)
		for j := range h.Keys {
			if h.Keys[j] == key {
				h.Keys = append(h.Keys[:j], h.Keys[j+1:]...)
			}
		}
	}
}

// GetImmediateNode will return the first node following the given node
// identified with its key. The node is found going counter clock-wise onto the ring
// In the case of servers, the Immediate node will represent the server to send 
// data to.
func (h *ConsistentHashing) GetImmediateNode(key string) string {
    h.Logger.Println("[INFO] GetImmediateNode ", key)

	if len(h.Nodes) == 0 {
		return ""
	}

	hash := h.computeHash(key)

	idx := sort.Search(len(h.Keys), func(i int) bool {
		return h.Keys[i] >= hash
	})

	if idx == len(h.Keys) {
		idx = 0
	}

	return h.Nodes[h.Keys[idx]]
}

// Private function not exported to be able to compute the hash of the provided key
func (h *ConsistentHashing) computeHash(uuid string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(uuid))
	return hash.Sum32()
}

// Hash hashes the given input using a graph-based data-structure and keeps a sorted list of nodes
// and Replicas for faster retrieval.
// It returns the immediate node index to which the uuid will be assigned
func (h *ConsistentHashing) Hash(uuid string, _ int) (string, error) {
	if len(uuid) == 0 {
	    h.Logger.Println("[ERROR] Consistent Hashing ", uuid, " failed")
		return "", errors.New("Expected uuid to be non-empty")
	}

	h.Logger.Println("[INFO] Hashing ", uuid, " succeeded")

	return h.GetImmediateNode(uuid), nil
}
