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

type ConsistentHashing struct {
	Nodes       map[uint32]string
	Replicas    int
	SortedNodes []uint32
	NodeIndex   map[string]int
}

func (h *ConsistentHashing) AddNode(node string) {
	for i := 0; i < h.Replicas; i++ {
		hash := h.computeHash(fmt.Sprintf("%s-%d", node, i))
		h.Nodes[hash] = node
		h.SortedNodes = append(h.SortedNodes, hash)
	}

	sort.Slice(h.SortedNodes, func(i, j int) bool {
		return h.SortedNodes[i] < h.SortedNodes[j]
	})

	if _, ok := h.NodeIndex[node]; !ok {
		h.NodeIndex[node] = len(h.NodeIndex)
	}
}

func (h *ConsistentHashing) RemoveNode(node string) {
	for i := 0; i < h.Replicas; i++ {
		hash := h.computeHash(fmt.Sprintf("%s-%d", node, i))
		delete(h.Nodes, hash)
		index := -1

		for j, v := range h.SortedNodes {
			if v == hash {
				index = j
				break
			}
		}

		if index != -1 {
			h.SortedNodes = append(h.SortedNodes[:index], h.SortedNodes[index+1:]...)
		}
	}
	delete(h.NodeIndex, node)
}

func (h *ConsistentHashing) GetImmediateNode(key string) int {
	if len(h.Nodes) == 0 {
		return -1
	}

	hash := h.computeHash(key)
	index := sort.Search(len(h.SortedNodes), func(i int) bool {
		return h.SortedNodes[i] >= hash
	})

	if index == len(h.SortedNodes) {
		index = 0
	}

	return h.NodeIndex[h.Nodes[h.SortedNodes[index]]]
}

func (h *ConsistentHashing) computeHash(uuid string) uint32 {
	hash := fnv.New32a()
	hash.Write([]byte(uuid))
	return hash.Sum32()
}

// Hash hashes the given input using a graph-based data-structure and keeps a sorted list of nodes
// and Replicas for faster retrieval.
// It returns the immediate node index to which the uuid will be assigned
func (h *ConsistentHashing) Hash(uuid string, replicas int) (int, error) {
    h.Replicas = replicas

	if len(uuid) == 0 {
		return -1, errors.New("Expected uuid to be non-empty")
	}

	return h.GetImmediateNode(uuid), nil
}

