// LRU 缓存机制 设计和实现一个 LRU（最近最少使用）缓存数据结构，使它应该支持一下操作：get 和 put。
// get(key) - 如果 key 存在于缓存中，则获取 key 的 value（总是正数），否则返回 -1。
// put(key,value) - 如果 key 不存在，请设置或插入 value。当缓存达到其容量时，它应该在插入新项目之前使最近最少使用的项目作废。
package main

import (
	"container/list"
	"fmt"
	"sync"
)

func main() {
	lru := LRU{}
	lru.Init(3)

	type kv struct {
		k string
		v int
	}
	kvs := []kv{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"b", 4},
		{"d", 5},
	}

	for _, kv := range kvs {
		lru.Set(kv.k, kv.v)
	}

	type testCase struct {
		input  string
		except int
	}

	cases := []testCase{
		{"a", -1},
		{"b", 4},
		{"c", 3},
		{"d", 5},
	}

	for _, cas := range cases {
		fmt.Println(lru.Get(cas.input) == cas.except)
	}
}

type LRU struct {
	sync.RWMutex
	capacity int
	lruInfo  *list.List
	dataPool map[string]*list.Element
}

type entry struct {
	key   string
	value int
}

func (it *LRU) Init(cap int) {
	it.capacity = cap
	it.lruInfo = list.New()
	it.dataPool = make(map[string]*list.Element, cap)
}

func (it *LRU) Get(key string) (value int) {
	it.RLock()
	// 检查是否存在该key
	element, ok := it.dataPool[key]
	it.RUnlock()

	// 如果不存在则返回-1
	if !ok {
		return -1
	}

	// 如果存在，则移动至队首
	it.Lock()
	it.lruInfo.MoveToFront(element)
	it.Unlock()

	// 返回数值
	return element.Value.(*entry).value
}

func (it *LRU) Set(key string, value int) {
	it.Lock()
	defer it.Unlock()

	// 检查当前是否存在该key
	element, ok := it.dataPool[key]

	if ok {
		// 如果key存在，则将其移动到队首
		it.lruInfo.MoveToFront(element)
		// 如果数值未变则直接返回
		if element.Value.(*entry).value == value {
			return
		}
		// 如果数值发生变动则进行更新
		element.Value.(*entry).value = value
		return
	}

	// 如果key不存在, 进行插入数值
	ent := &entry{
		key:   key,
		value: value,
	}
	front := it.lruInfo.PushFront(ent)
	it.dataPool[key] = front

	// 如果LRU的容量在限制内则直接返回
	if len(it.dataPool) <= it.capacity {
		return
	}

	// 如果当前容量超过最大值，则删除最旧的key
	back := it.lruInfo.Back()
	it.lruInfo.Remove(back)
	oldKey := back.Value.(*entry).key
	delete(it.dataPool, oldKey)
	return
}
