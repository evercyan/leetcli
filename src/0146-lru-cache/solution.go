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

func (this *LRUCache) Get(key int) int {
	val, ok := this.Hash[key]
	if !ok {
		return -1
	}
	this.MoveToFront(key)
	return val
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.Hash[key]; !ok {
		// 添加
		if this.Capacity <= len(this.Hash) {
			// 如果长度超过了, 删除链表中最后位, 同时清除其 hash 数据
			lastKey := this.List.Back().Value
			delete(this.Hash, lastKey.(int))
			this.List.Remove(this.List.Back())
		}
		this.Hash[key] = value
		this.List.PushFront(key)
	} else {
		// 更新
		this.Hash[key] = value
		this.MoveToFront(key)
	}
	return
}

func (this *LRUCache) MoveToFront(key int) {
	for v := this.List.Front(); v != nil; v = v.Next() {
		if key == v.Value {
			this.List.MoveToFront(v)
		}
	}
	return
}
