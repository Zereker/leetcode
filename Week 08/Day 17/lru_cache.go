package main

import (
	"container/list"
)

type LRUCache struct {
	cap int                   // capacity
	l   *list.List            // doubly linked list
	m   map[int]*list.Element // hash table for checking if list node exists
}

// Pair is the value of a list node.
type Pair struct {
	key   int
	value int
}

// Constructor initializes a new LRUCache.
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   list.New(),
		m:   make(map[int]*list.Element, capacity),
	}
}

// Get a list node from the hash map.
func (c *LRUCache) Get(key int) int {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		val := node.Value.(Pair).value
		// move node to front
		c.l.MoveToFront(node)
		return val
	}

	return -1
}

// Put key and value in the LRUCache
func (c *LRUCache) Put(key int, value int) {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		// move the node to front
		c.l.MoveToFront(node)
		// update the value of a list node
		node.Value = Pair{key: key, value: value}
	} else {
		// delete the last list node if the list is full
		if c.l.Len() == c.cap {
			// get the key that we want to delete
			idx := c.l.Back().Value.(Pair).key
			// delete the node pointer in the hash map by key
			delete(c.m, idx)
			// remove the last list node
			c.l.Remove(c.l.Back())
		}
		// initialize a list node
		node := Pair{key: key, value: value}

		// push the new list node into the list
		ptr := c.l.PushFront(node)
		// save the node pointer in the hash map
		c.m[key] = ptr
	}
}

func main() {

	cache := Constructor(2)

	cache.Put(1, 1)
	cache.Put(2, 2)
	println(cache.Get(1)) // 返回  1
	cache.Put(3, 3)       // 该操作会使得关键字 2 作废
	println(cache.Get(2)) // 返回 -1 (未找到)
	cache.Put(4, 4)       // 该操作会使得关键字 1 作废
	println(cache.Get(1)) // 返回 -1 (未找到)
	println(cache.Get(3)) // 返回  3
	println(cache.Get(4)) // 返回  4
}
