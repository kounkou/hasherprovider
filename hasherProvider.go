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

package hasherProvider

import (
	"fmt"

	consistent "github.com/kounkou/hasherProvider/consistent"
	random "github.com/kounkou/hasherProvider/random"
	uniform "github.com/kounkou/hasherProvider/uniform"
)

const (
	CONSISTENT_HASHING = 0
	RANDOM_HASHING     = 1
	UNIFORM_HASHING    = 2
)

type Hasher interface {
    Hash(uuid string, n int) (int, error)
    AddNode(uuid string)
    RemoveNode(uuid string)
}

type HasherProvider struct {
	hasherMap map[int]Hasher
}

func (h *HasherProvider) GetHasher(hashFunction int) (Hasher, error) {
	h.initHasherMap()
	hasher, ok := h.hasherMap[hashFunction]

	if !ok {
		return nil, fmt.Errorf("unknown hashing function type: %d", hashFunction)
	}
	return hasher, nil
}

func (h *HasherProvider) initHasherMap() {
	h.hasherMap = map[int]Hasher{
		CONSISTENT_HASHING: &consistent.ConsistentHashing{},
		RANDOM_HASHING:     &random.RandomHashing{},
		UNIFORM_HASHING:    &uniform.UniformHashing{},
	}
}
