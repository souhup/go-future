/*
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

进阶:

你是否可以在 O(1) 时间复杂度内完成这两种操作？

示例:

LRUCache cache = new LRUCache( 2 缓存容量 );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1); // 返回  1
cache.put(3, 3); // 该操作会使得密钥 2 作废
cache.get(2); // 返回 -1 (未找到)
cache.put(4, 4); // 该操作会使得密钥 1 作废
cache.get(1);    // 返回 -1 (未找到)
cache.get(3); // 返回  3
cache.get(4); // 返回  4
*/

package main

import "container/list"

func main() {
	var cache = Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	println(cache.Get(1))
	cache.Put(3, 3)
	println(cache.Get(2))
	cache.Put(4, 4)
	println(cache.Get(1))
	println(cache.Get(3))
	println(cache.Get(4))
}

type Entry struct {
	key     int
	value   int
	pointer *list.Element
}

type LRUCache struct {
	capacity  int
	dictCache map[int]*Entry
	linkCache *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity:  capacity,
		dictCache: make(map[int]*Entry),
		linkCache: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	value, ok := this.dictCache[key]
	if !ok {
		return -1
	}
	this.linkCache.MoveToFront(value.pointer)
	return value.value
}

func (this *LRUCache) Put(key int, value int) {
	v, ok := this.dictCache[key]
	if !ok {
		if len(this.dictCache) == this.capacity {
			pointer := this.linkCache.Back()
			delete(this.dictCache, pointer.Value.(int))
			this.linkCache.Remove(pointer)
		}
		pointer := this.linkCache.PushFront(key)
		entry := &Entry{
			key:     key,
			value:   value,
			pointer: pointer,
		}
		this.dictCache[key] = entry
	} else {
		this.linkCache.Remove(v.pointer)
		pointer := this.linkCache.PushFront(key)
		entry := &Entry{
			key:     key,
			value:   value,
			pointer: pointer,
		}
		this.dictCache[key] = entry
	}
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
