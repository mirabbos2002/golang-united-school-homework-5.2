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
	v, ok := c.c[key]
	if ok {
		if time.Now().After(v.time) && (v.time != time.Time{}) {
			delete(c.c, key)
			return "", false
		} else {
			return v.value, true
		}
	}
	return "", false
}

func (c *Cache) Put(key, value string) {
	c.c[key] = valueTime{value: value}
}

func (c *Cache) Keys() (a []string) {
	for k, v := range c.c {
		if time.Now().After(v.time) && (v.time != time.Time{}) {
			delete(c.c, k)
		} else {
			a = append(a, k)
		}
	}
	return
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.c[key] = valueTime{value, deadline}
}
