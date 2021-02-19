// Package lru implements a simple lru cache.
package lru

import (
	"container/list"
)

type LRUCache struct {
	// maxCap sets the max elements the cache supports. 0 or negative number means no limit
	maxCap          int
	ll              *list.List
	cache           map[interface{}]*list.Element
	evictedCallback func(key, value interface{})
	// size is the number of element in this cache.
	size int
}

type entry struct {
	key   interface{}
	value interface{}
}

// NewCache returns a LRU cache with fixed size. If 0 or negative number, then there is no limit.
func NewCache(maxCap int) *LRUCache {
	if maxCap < 0 {
		maxCap = 0
	}
	return &LRUCache{
		maxCap: maxCap,
		ll:     list.New(),
		cache:  make(map[interface{}]*list.Element),
	}
}

// SetEvictedCallback set the evicated callback func for current cache
func (l *LRUCache) SetEvictedCallback(evicted func(key, value interface{})) {
	l.evictedCallback = evicted
}

// Add adds a new key/value into the cache.
func (l *LRUCache) Add(key, value interface{}) {
	if l.cache == nil {
		l.cache = make(map[interface{}]*list.Element)
		l.ll = list.New()
	}

	if ele, ok := l.cache[key]; ok {
		l.ll.MoveToFront(ele)
		ele.Value.(*entry).value = value
		return
	}
	l.cache[key] = l.ll.PushFront(&entry{key: key, value: value})
	l.size += 1
	if l.maxCap != 0 && l.size > l.maxCap {
		l.removeOldest()
		l.size -= 1
	}
}

// Get returns the value based on the key.
func (l *LRUCache) Get(key interface{}) (interface{}, bool) {
	if l.cache == nil {
		return nil, false
	}

	if ele, ok := l.cache[key]; ok {
		l.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}

	return nil, false
}

// Remove deletes a element from the cache.
func (l *LRUCache) Remove(key interface{}) {
	if l.cache == nil {
		return
	}

	if ele, ok := l.cache[key]; ok {
		l.removeElement(ele)
	}
}

func (l *LRUCache) removeOldest() {
	if l.cache == nil {
		return
	}
	ele := l.ll.Back()
	if ele != nil {
		l.removeElement(ele)
	}
}

func (l *LRUCache) removeElement(ele *list.Element) {
	l.ll.Remove(ele)
	e := ele.Value.(*entry)
	delete(l.cache, e.key)
	l.size -= 1
	if l.evictedCallback != nil {
		l.evictedCallback(e.key, e.value)
	}
}

// Len returns the number of element in the cache.
func (l *LRUCache) Len() int {
	return l.size
}

// Purge clears the cache
func (l *LRUCache) Purge() {
	if l.evictedCallback != nil {
		for _, ele := range l.cache {
			e := ele.Value.(*entry)
			l.evictedCallback(e.key, e.value)
		}
	}
	l.cache = nil
	l.ll = nil
	l.size = 0
}
