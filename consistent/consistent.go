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
)

// With consistent Hashing, the keys already assigned to a shard
// do NOT need to be reassigned. Hence solving the issue introduced
// by the usage of Modulo to be able to perform a consistent Hashing.

type ConsistentHashing struct {
	Nodes       map[uint32]string
	Replicas    int
	Keys        []uint32
}

func (h *ConsistentHashing) SetReplicas(replicas int) {
    h.Replicas = replicas
}

func (h *ConsistentHashing) AddNode(node string) {
	for i := 0; i < h.Replicas; i++ {
		key := h.computeHash(node + strconv.Itoa(i))
		h.Nodes[key] = node
		h.Keys = append(h.Keys, key)
	}

	sort.Slice(h.Keys, func(i, j int) bool { return h.Keys[i] < h.Keys[j] })
}

func (h *ConsistentHashing) RemoveNode(node string) {
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

func (h *ConsistentHashing) GetImmediateNode(key string) string {
	if len(h.Nodes) == 0 {
		return "NA"
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

func (h *ConsistentHashing) computeHash(uuid string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(uuid))
	return hash.Sum32()
}

// Hash hashes the given input using a graph-based data-structure and keeps a sorted list of nodes
// and Replicas for faster retrieval.
// It returns the immediate node index to which the uuid will be assigned
// TODO : Change parameter list to be a struct so that we can ignore second parameter
func (h *ConsistentHashing) Hash(uuid string, _ int) (string, error) {
	if len(uuid) == 0 {
		return "NA", errors.New("Expected uuid to be non-empty")
	}

	return h.GetImmediateNode(uuid), nil
}

