package Singleton

import (
	"sync"
	"sync/atomic"
)

type SingleTon_1 struct {
}

var (
	instance_1  *SingleTon_1
	initialized uint32
	mu          sync.Mutex
)

func Instance_1() *SingleTon_1 {
	if 1 == atomic.LoadUint32(&initialized) {
		return instance_1
	}

	mu.Lock()
	mu.Unlock()
	if nil == instance_1 {
		defer atomic.StoreUint32(&initialized, 1)
		instance_1 = new(SingleTon_1)
	}
	return instance_1
}
