[![Go](https://github.com/kounkou/hasherprovider/workflows/Go/badge.svg)](https://github.com/kounkou/hasherprovider/actions?query=workflow%3AGo)
[![Coverage Status](https://coveralls.io/repos/github/kounkou/hasherprovider/badge.svg?branch=main)](https://coveralls.io/github/kounkou/hasherprovider?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/kounkou/hasherprovider)](https://goreportcard.com/report/github.com/kounkou/hasherprovider)
[![gopkg](https://pkg.go.dev/badge/github.com/kounkou/hasherprovider.svg)](https://pkg.go.dev/github.com/kounkou/hasherprovider)
[![license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/kounkou/hasherprovider/blob/master/LICENSE)

# hasherprovider

The Hasher library implements 3 hashing algorithms (Consistent, Uniform and Random) on a given key string or UUID and returns the index (for Uniform and Random algorithms), and the string of the node (for Consistent algorithm) to which the given string should be mapped.
Consistent hashing is one such algorithm that minimizes the number of updates required to associate the request with the appropriate server. 
This addresses the common problem of reassigning servers that arises when using the modulo operation.

# Installation

```bash
go get github.com/kounkou/hasherprovider
```

# Usage

```golang
package main

import (
    "fmt"
    "github.com/kounkou/hasherprovider"
)

const (
	CONSISTENT_HASHING = 0
	RANDOM_HASHING     = 1
	UNIFORM_HASHING    = 2
)

func main() {
	// Create a new HasherProvider object
	hasherprovider := hasherprovider.HasherProvider{ }

	// Get the consistent hashing function
	h, err := hasherprovider.GetHasher(CONSISTENT_HASHING)

    	// Set replicas entities
	h.SetReplicas(1)

	if h == nil || err != nil {
		fmt.Println("Error getting hasher:", err)
		return
	}

    	h.AddNode("server1")
	h.AddNode("server2")
    	h.AddNode("server3")

    	result, err := h.Hash("9", 0)

    	if err != nil {
        	fmt.Println("Error getting hash for some string `9` ", err)
    	}

    	if result != "server2" {
        	fmt.Errorf("Expected replica to be assigned is server2")
        	return
    	}

    	fmt.Println("Success... !")
}
```

# Algorithms

HasherProvider currently supports 3 algorithms. You might want to choose your hashing algorithm based on the following characteristics :

| Hashing Algorithm  | Load balanced | Elastic   | Fault tolerant | Decentralized |
|--------------------|---------------|-----------|----------------|---------------|
| Consistent Hashing | Excellent     | Excellent | Excellent      | Excellent     |
| Random Hashing     | Good          | Poor      | Good           | Poor          |
| Uniform Hashing    | Poor          | Good      | Poor           | Excellent     |

