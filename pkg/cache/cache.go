package cache

import {
	"dll"
}

type LRUCache {
	capacity int
	dict map[int][*Node]
	dll *DLL
}

func (lru *LRUCache) Put(key int, value int) {
	node, contains := lru.dict[key]
	if contains {
		lru.dll.DeleteNode(node)
		lru.dll.AddToHead(node)
		lru.dict[key].value = value
	} else {
		if len(lru.dict) == capacity {
			lru.dll.RemoveFromTail()
			delete(lru.dict, key)
		} else {
			newNode := &Node{key: key, value: value}
			lru.dll.AddToHead(newNode)
			lru.dict[key] = newNode
		}
	}
}

func (lru *LRUCache) Get(key int) int {
	node, contains := lru.dict[key]
	if contains {
		lru.dll.DeleteNode(node)
		lru.dll.AddToHead(node)
		return node.value
	} else {
		return -1
	}
}