# in-memory-cache

## Description
It's a simple in-memory cache written using Go. Cache constructor has 1 attribute it's a time interval the cache will start cleanup process

## Example
```go
package main

import (
	"fmt"

	cache "github.com/papvan/in-memory-cache"
)

func main() {

	cache := cache.New(time.Second * 5)

	cache.Set("userId", 42, time.Second * 7)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId := cache.Get("userId")

	fmt.Println(userId)
}
```