package utils

import "sync"

// Observable 观察者工具类：用于并发控制
type Observable struct {
	mu    sync.Mutex
	value any
	cond  *sync.Cond
}

func NewObservable(initial any) *Observable {
	o := &Observable{value: initial}
	o.cond = sync.NewCond(&o.mu)
	return o
}

func (o *Observable) Set(v any) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.value = v
	o.cond.Broadcast() // 通知所有等待的监听者
}

func (o *Observable) Get() any {
	o.mu.Lock()
	defer o.mu.Unlock()
	return o.value
}

func (o *Observable) Wait() {
	o.cond.Wait()
}

func (o *Observable) Lock() {
	o.mu.Lock()
}

func (o *Observable) UnLock() {
	o.mu.Unlock()
}

func (o *Observable) Value() any {
	return o.value
}
