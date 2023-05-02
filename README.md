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
    uuid "github.com/google/uuid"
    hash "github.com/kounkou/hasherProvider"
)

const (
    CONSISTENT_HASHING = 0
    RANDOM_HASHING     = 1
    UNIFORM_HASHING    = 2
)

func main() {

    hasherProvider := &hash.HasherProvider{}
    hasher, err := hasherProvider.GetHasher(RANDOM_HASHING)

    if err != nil {
        fmt.Println("Handle error ", err)
    }

    event := uuid.New().String()
    result, err := hasher.Hash(event, 10)

    if err != nil {
        fmt.Println("Handle error ", err)
    }

    fmt.Println("Hashing result  ", result)
}
```

# Algorithms

HasherProvider currently supports 3 algorithms : 

- Consistent Hashing
- Random Hashing
- Uniform Hashing

