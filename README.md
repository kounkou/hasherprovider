# hasherProvider

Hasher library that applies hashing on a given key string or uuid.

# Installation

```bash
go get github.com/kounkou/hashProvider
```

# Usage

```golang
import (
  hash "github.com/kounkou/hasherProvider"
  uuid "github.com/google/uuid"
  "fmt"
)

const (
	CONSISTENT_HASHING = 0
	RANDOM_HASHING     = 1
	UNIFORM_HASHING    = 2
)
  
hasherProvider := hash.HasherProvider{}
hasher, err := hasherProvider.GetHasher(RANDOM_HASHING)

if err != nil {
    fmt.Println("Handle error ", err)
}

event := uuid.New().String()

fmt.Println("Hashing result  ", hasher.Hash(event))
```

# Algorithms

HasherProvider currently supports 3 algorithms : 

- Consistent Hashing
- Random Hashing
- Uniform Hashing

