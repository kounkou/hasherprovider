# hasherProvider

Hasher library that applies hashing on a given key string or uuid.

# Installation

```bash
go get github.com/google/uuid
go get github.com/kounkou/hasherProvider
```

# Usage

```golang
package main

import (
    "fmt"
	"github.com/kounkou/hasherProvider"
)

const (
	CONSISTENT_HASHING = 0
	RANDOM_HASHING     = 1
	UNIFORM_HASHING    = 2
)

func main() {
	// Create a new HasherProvider object
	hasherProvider := hasherProvider.HasherProvider{ }

	// Get the consistent hashing function
	h, err := hasherProvider.GetHasher(CONSISTENT_HASHING)

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
        fmt.Println("Error getting hash for `node2` ", err)
    }

    if result != "server2" {
        fmt.Errorf("Expected replica to be assigned is server1")
        return
    }

    fmt.Println("Success... !")
}
```

# Algorithms

HasherProvider currently supports 3 algorithms : 

- Consistent Hashing
- Random Hashing
- Uniform Hashing

