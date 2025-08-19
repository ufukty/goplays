package cache

import "fmt"

type Cache struct{}

func New() *Cache {
	fmt.Println("Constructing Cache")
	return &Cache{}
}
