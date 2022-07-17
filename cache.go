package cache

import (
	"time"
)

type valueTime struct {
	value string
	time  time.Time
}
type Cache struct {
	c map[string]valueTime
}

// var wg sync.WaitGroup

func NewCache() *Cache {
	return &Cache{map[string]valueTime{}}
}

func (c Cache) Get(key string) (string, bool) {
	if v := c.c[key]; time.Now().After(v.time) {
		delete(c.c, key)
		return "", false
	} else {
		defer delete(c.c, key)
		return c.c[key].value, true
	}

	// if k, ok := c.c[key]; ok {

	// 	emp := Cache{}
	// 	if k.deadline == emp.deadline {
	// 		return caches[key].value, true
	// 	}else if time.Now().After(k.deadline) {
	// 		delete(caches, key)
	// 		return "", false
	// 	} else {
	// 		return caches[key].value, true
	// 	}
	// }
	// return "", false
}

func (c *Cache) Put(key, value string) {
	c.c[key] = valueTime{value: value}
}

func (c *Cache) Keys() (a []string) {
	for k, v := range c.c {
		if time.Now().After(v.time) {
			delete(c.c, k)
		} else {
			a = append(a, v.value)
		}
	}
	return
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.c[key] = valueTime{value, deadline}
}
