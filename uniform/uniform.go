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

package uniform

import (
	"errors"
)

type Hasher interface {
    Hash(uuid string, n int) (int, error)
    AddNode(uuid string)
    RemoveNode(uuid string)
}

type UniformHashing struct {
	values []int
}

// Uniform hashing is used to distribute the uuid's associated (example events...)
// across the a set of shards indexed.
// Uniform hashing makes sense when the number of shards is fixed. For dynamic shards
// please consider using `consistent hashing`
func (h UniformHashing) Hash(uuid string, shards int) (int, error) {
	if shards == 0 || len(uuid) == 0 {
		return 0, errors.New("Expected shards to be positive non 0")
	}

	hash := 0
	for i := 0; i < len(uuid); i++ {
		hash = (hash << 5) + hash + int(uuid[i])
	}
	return hash % shards, nil
}

// Implemented for convenience
func (h UniformHashing) AddNode(_ string) {
}
// Implemented for convenience
func (h UniformHashing) RemoveNode(_ string) {
}

