# in-memory-cache

## Description
It's a simple in-memory cache written using Go

## Example
```go
package main

import (
	"fmt"

	cache "github.com/papvan/in-memory-cache"
)

func main() {

	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}
```