package solution

import (
	"container/list"
)

type LRUCache struct {
	Capacity int
	Hash     map[int]int
	List     *list.List
}

func Constructor(capacity int) LRUCache {
	r := LRUCache{
		Capacity: capacity,
		Hash:     map[int]int{},
		List:     list.New(),
	}
	return r
}

func (c *LRUCache) Get(key int) int {
	val, ok := c.Hash[key]
	if !ok {
		return -1
	}
	c.MoveToFront(key)
	return val
}

func (c *LRUCache) Put(key int, value int) {
	if _, ok := c.Hash[key]; !ok {
		// 添加
		if c.Capacity <= len(c.Hash) {
			// 如果长度超过了, 删除链表中最后位, 同时清除其 hash 数据
			lastKey := c.List.Back().Value
			delete(c.Hash, lastKey.(int))
			c.List.Remove(c.List.Back())
		}
		c.Hash[key] = value
		c.List.PushFront(key)
	} else {
		// 更新
		c.Hash[key] = value
		c.MoveToFront(key)
	}
	return
}

func (c *LRUCache) MoveToFront(key int) {
	for v := c.List.Front(); v != nil; v = v.Next() {
		if key == v.Value {
			c.List.MoveToFront(v)
		}
	}
	return
}
