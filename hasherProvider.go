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
	"log"
	"os"

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
	Hash(uuid string, n int) (string, error)
	AddNode(uuid string)
	RemoveNode(uuid string)
	SetReplicas(replicas int)
}

type HasherProvider struct {
    Logger *log.Logger
}

func (h *HasherProvider) GetHasher(hashFunction int) (Hasher, error) {
    if h.Logger == nil {
        h.Logger = log.New(os.Stdout, "hasherProvider ", log.LstdFlags)
        h.Logger.Println("[WARN] Setting default logger")
    }

    h.Logger.Println("[INFO] Getting Hasher with hashing algorithm ", hashFunction)

	hasherMap := h.initHasherMap()
	hasher, ok := hasherMap[hashFunction]

	if !ok {
	    h.Logger.Println("[ERROR] Getting the hasher failed for ", hashFunction)
		return nil, fmt.Errorf("unknown hashing function type: %d", hashFunction)
	}

	h.Logger.Println("[INFO] Getting Hasher with hashing algorithm ", hashFunction, " succeeded")

	return hasher, nil
}

func (h *HasherProvider) initHasherMap() map[int]Hasher {
    h.Logger.Println("[INFO] InitHasherMap")

	var hasherMap map[int]Hasher

	hasherMap = map[int]Hasher{
		CONSISTENT_HASHING: &consistent.ConsistentHashing{
			Nodes:    make(map[uint32]string),
			Replicas: 0,
			Logger:   h.Logger,
		},
		RANDOM_HASHING:  &random.RandomHashing{},
		UNIFORM_HASHING: &uniform.UniformHashing{},
	}

	h.Logger.Println("[INFO] InitHasherMap successfully")

	return hasherMap
}
