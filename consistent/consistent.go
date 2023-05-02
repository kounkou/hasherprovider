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
	"fmt"
	"hash/fnv"
	"sort"
)

// With consistent Hashing, the keys already assigned to a shard
// do NOT need to be reassigned. Hence solving the issue introduced
// by the usage of Modulo to be able to perform a consistent Hashing.

type Hasher interface {
	Hash(uuid string, n int) (string, error)
}

type ConsistentHashing struct {
	nodes       map[uint32]string
	replicas    int
	sortedNodes []uint32
}

func (h *ConsistentHashing) AddNode(node string) {
	for i := 0; i < h.replicas; i++ {
		hash := h.computeHash(fmt.Sprintf("%s-%d", node, i))
		h.nodes[hash] = node
		h.sortedNodes = append(h.sortedNodes, hash)
	}

	sort.Slice(h.sortedNodes, func(i, j int) bool {
		return h.sortedNodes[i] < h.sortedNodes[j]
	})
}

func (h *ConsistentHashing) RemoveNode(node string) {
	for i := 0; i < h.replicas; i++ {
		hash := h.computeHash(fmt.Sprintf("%s-%d", node, i))
		delete(h.nodes, hash)
		index := 1

		for j, v := range h.sortedNodes {
			if v == hash {
				index = j
				break
			}
		}

		if index != -1 {
			h.sortedNodes = append(h.sortedNodes[:index], h.sortedNodes[index+1:]...)
		}
	}
}

func (h *ConsistentHashing) GetImmediateNode(key string) string {
	if len(h.nodes) == 0 {
		return ""
	}

	hash := h.computeHash(key)
	index := sort.Search(len(h.sortedNodes), func(i int) bool {
		return h.sortedNodes[i] >= hash
	})

	if index == len(h.sortedNodes) {
		index = 0
	}

	return h.nodes[h.sortedNodes[index]]
}

func (h *ConsistentHashing) computeHash(uuid string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(uuid))
	return hash.Sum32()
}

// Hash hashes the given input using a graph-based data-structure and keeps a sorted list of nodes
// and replicas for faster retrieval.
// It returns the immediate node to which the uuid will be assigned
func (h *ConsistentHashing) Hash(uuid string) (string, error) {
	if len(uuid) == 0 {
		return "", errors.New("Expected uuid to be non-empty")
	}

	return h.GetImmediateNode(uuid), nil
}
