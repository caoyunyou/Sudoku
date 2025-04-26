package event

import (
	"sync"
	"sync/atomic"
)

type Event struct {
	Type string
	Data interface{}
}

// Handler 事件处理器
type Handler func(event Event)

type Subscription struct {
	Cancel func() // 取消订阅的函数
}

// Bus 事件总线
type Bus struct {
	handlers map[string][]*handlerEntry
	mu       sync.RWMutex
}

// NewEventBus 创建事件总线
func NewEventBus() *Bus {
	return &Bus{
		handlers: make(map[string][]*handlerEntry),
	}
}

// handlerEntry 处理器
type handlerEntry struct {
	handler Handler
	id      int64 // 唯一标识符
}

var (
	nextID int64 // 自增 ID 生成器
)

// Subscribe 事件订阅:返回取消函数
func (eb *Bus) Subscribe(eventType string, handler Handler) func() {
	id := atomic.AddInt64(&nextID, 1)
	return eb.subscribeWithCancel(eventType, handler, id)
}

// subscribeWithCancel 实际订阅逻辑
func (eb *Bus) subscribeWithCancel(eventType string, handler Handler, id int64) func() {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	entry := &handlerEntry{
		handler: handler,
		id:      id,
	}

	if eb.handlers == nil {
		eb.handlers = make(map[string][]*handlerEntry)
	}
	eb.handlers[eventType] = append(eb.handlers[eventType], entry)

	return func() {
		eb.unsubscribe(eventType, id)
	}
}

// unsubscribe 取消订阅：通过唯一标识符
func (eb *Bus) unsubscribe(eventType string, id int64) {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	handlers := eb.handlers[eventType]
	if handlers == nil {
		return
	}

	// 过滤掉指定 ID 的处理函数
	var newHandlers []*handlerEntry
	for _, entry := range handlers {
		if entry.id != id {
			newHandlers = append(newHandlers, entry)
		}
	}
	eb.handlers[eventType] = newHandlers
}

// Publish 发布事件
func (eb *Bus) Publish(event Event) {
	eb.mu.RLock()
	handlersList, exists := eb.handlers[event.Type]
	eb.mu.RUnlock()

	if !exists {
		return
	}

	for _, entry := range handlersList {
		go entry.handler(event)
	}
}
