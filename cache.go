package cache

import (
	"sync"
	"time"
)

type Cache struct {
	value    string
	deadline time.Time
}

var wg sync.WaitGroup

func NewCache() *Cache {
	return &Cache{}
}

var caches = map[string]Cache{}

func (c *Cache) Get(key string) (string, bool) {
	if k, ok := caches[key]; ok {
		emp := Cache{}
		if k.deadline == emp.deadline {
			return caches[key].value, true
		}else if time.Now().After(k.deadline) {
			delete(caches, key)
			return "", false
		} else {
			return caches[key].value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	caches[key] = Cache{value: value}
}

func (c *Cache) Keys() (a []string) {
	for k, v := range caches {
		if time.Now().After(v.deadline) {
			delete(caches, k)
		} else {
			a = append(a, v.value)
		}
	}
	return
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	caches[key] = Cache{value, deadline}
}
