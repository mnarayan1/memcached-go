package lru

import (
	dll "dll"
)

type LRUCache struct {
	Capacity int
	Cache map[string]*Node
	List *DLL
}