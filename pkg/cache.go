package gocache

import (
	"dll"
	"sync"
	"time"
)

type Cache struct {
	capacity int
	dict     map[string]*dll.Node
	dll      *dll.DLL
	mutex    sync.Mutex
}

func NewCache(capacity int) *Cache {
	newDLL := dll.DLLInit()
	return &Cache{
		capacity: capacity,
		dict:     make(map[string]*dll.Node),
		dll:      newDLL,
	}
}

func (lru *Cache) Set(key string, value string, ttl time.Duration) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	node, contains := lru.dict[key]
	if contains {
		lru.dll.DeleteNode(node)
		node.value = value
		node.expiration = time.Now().Add(ttl)
		lru.dll.AddToHead(node)
	} else {
		if len(lru.dict) == lru.capacity {
			toRemove := lru.dll.RemoveFromTail()
			if toRemove != nil {
				delete(lru.dict, toRemove.key)
			}
		}
		newNode := &dll.Node{
			key:        key,
			value:      value,
			expiration: time.Now().Add(ttl),
		}
		lru.dll.AddToHead(newNode)
		lru.dict[key] = newNode
	}
}

func (lru *Cache) Get(key string) string {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	node, contains := lru.dict[key]
	if contains {
		if time.Now().After(node.expiration) {
			lru.dll.DeleteNode(node)
			delete(lru.dict, key)
			return "Not found"
		}
		lru.dll.DeleteNode(node)
		lru.dll.AddToHead(node)
		return node.value
	}
	return "Not found"
}

func (lru *Cache) Delete(key string) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	node, contains := lru.dict[key]
	if contains {
		lru.dll.DeleteNode(node)
		delete(lru.dict, key)
	}
}