package cache

import {
	"dll",
	"sync"
}

type LRUCache struct {
	capacity int
	dict map[string]*dll.Node
	dll *dll.DLL
	mutex sync.Mutex
}

func (lru *LRUCache) Put(key string, value string) {
	lru.mutex.Lock()
	defer lru.mutex.Unlock()
	
	node, contains := lru.dict[key]
	if contains {
		lru.dll.DeleteNode(node)
		node.value = value
		lru.dll.AddToHead(node)
	} else {
		if len(lru.dict) == lru.capacity {
			toRemove := lru.dll.RemoveFromTail()
			delete(lru.dict, toRemove.key)
		}
		newNode := &dll.Node{key: key, value: value}
		lru.dll.AddToHead(newNode)
		lru.dict[key] = newNode
	}
}

func (lru *LRUCache) Get(key string) string {
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