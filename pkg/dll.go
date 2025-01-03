package dll

import (
	"sync"
	"time"
)

type Node struct {
	key        string
	value      string
	prev       *Node
	next       *Node
	expiration time.Time
}

// dll with sentinel head/tail
type DLL struct {
	head *Node
	tail *Node
	mutex sync.Mutex
}

func DLLInit() *DLL {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return &DLL{head: head, tail: tail}
}

func (dll *DLL) AddToHead(node *Node) {
	dll.mutex.Lock()
	defer dll.mutex.Unlock()

	oldHead := dll.head.next
	oldHead.prev = node
	node.next = oldHead
	node.prev = dll.head
	dll.head.next = node
}

func (dll *DLL) RemoveFromTail() *Node {
	dll.mutex.Lock()
	defer dll.mutex.Unlock()

	if dll.tail.prev != dll.head {
		toRemove := dll.tail.prev
		dll.DeleteNode(toRemove)
		return toRemove
	}
	return nil
}

func (dll *DLL) DeleteNode(node *Node) {
	dll.mutex.Lock()
	defer dll.mutex.Unlock()

	node.prev.next = node.next
	node.next.prev = node.prev
}