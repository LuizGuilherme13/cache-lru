package main

import (
	"sort"
	"time"
)

type Cache struct {
	capacity int
	items    []Item
}

type Item struct {
	key    string
	value  any
	access time.Time
}

func New(capacity int) *Cache {
	return &Cache{
		capacity: capacity,
	}
}

func (l *Cache) Get(key string) any {
	for i := range l.items {
		if l.items[i].key == key {
			l.items[i].access = time.Now()

			sort.Slice(l.items, func(i, j int) bool {
				return l.items[i].access.Before(l.items[j].access)
			})

			return l.items[i].value
		}
	}

	return -1
}

func (l *Cache) Set(key string, value any) {
	if l.capacity == len(l.items) {
		l.items[0].key = key
		l.items[0].value = value
		l.items[0].access = time.Time{}
		return
	}

	l.items = append(l.items, Item{key: key, value: value})
}
