package gocache

import {
	"dll",
	"sync"
}

// cache with lru eviction

type Cache struct {
	capacity int
	dict map[string]*dll.Node
	dll *dll.DLL
	mutex sync.Mutex
}

func () NewCache(capacity int) *Cache{
	newDLL := dll.DLLInit()
	newCache := &Cache{capacity: capacity, dll: newDLL}
}

func (lru *Cache) Set(key string, value string) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	
	node, contains := lru.dict[key]
	if contains {
		lru.dll.DeleteNode(node)
		node.value = value
		lru.dll.AddToHead(node)
	} else {
		if len(lru.dict) == lru.capacity {
			toRemove := lru.dll.tail
			lru.dll.RemoveFromTail()
			delete(lru.dict, toRemove.key)
		}
		newNode := &dll.Node{key: key, value: value}
		lru.dll.AddToHead(newNode)
		lru.dict[key] = newNode
	}
}

func (lru *Cache) Get(key string) string {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()

	node, contains := lru.dict[key]
	if contains {
		lru.dll.DeleteNode(node)
		lru.dll.AddToHead(node)
		return node.value
	}
	return "Not found"
}

func (lru *Cache) Delete(key string) {
	node, contains := lru.dict[key]
	if contains {
		lru.dll.RemoveFromTail()
		delete(lru.dict, node.key)
	}
}