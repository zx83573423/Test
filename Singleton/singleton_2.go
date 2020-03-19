package Singleton

import "sync"

type Singleton_2 struct {
}

var (
	singleton_2 *Singleton_2
	once        sync.Once
)

func Instance_2() *Singleton_2 {
	once.Do(func() {
		singleton_2 = new(Singleton_2)
	})
	return singleton_2
}
