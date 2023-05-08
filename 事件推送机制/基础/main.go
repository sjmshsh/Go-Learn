// 事件源：你需要明确事件源，即哪些事件需要被推送出去。
// 在 etcd v3 中，事件源包括 key 的创建、修改、删除等操作。你需要确定你的事件源，并
// 为每个事件源定义事件类型和事件数据结构。

// 事件订阅：你需要提供一个订阅机制，让客户端能够订阅感兴趣的事件。
// 在 etcd v3 中，客户端可以通过 gRPC API 订阅感兴趣的事件。你需要定义类似的 API，
// 让客户端能够订阅和取消订阅事件。

// 事件推送：你需要实现事件推送逻辑，即当事件发生时，将事件推送给订阅者。
// 你可以采用类似于观察者模式的设计，即为每个事件源维护一个订阅者列表，当事件发生时，依次通知订阅者。

package main

import (
	"fmt"
	"sync"
)

type Event struct {
	Type string
	Data interface{}
}

type Subscriber func(event Event)

type EventSource struct {
	mu          sync.RWMutex
	subscribers []Subscriber
}

func (es *EventSource) Subscribe(subscriber Subscriber) {
	es.mu.Lock()
	defer es.mu.Unlock()
	es.subscribers = append(es.subscribers, subscriber)
}

//func (es *EventSource) Unsubscribe(subscriber Subscriber) {
//	es.mu.Lock()
//	defer es.mu.Unlock()
//	for i, s := range es.subscribers {
//		if s == subscriber {
//			es.subscribers = append(es.subscribers[:i], es.subscribers[i+1:]...)
//			return
//		}
//	}
//}

func (es *EventSource) Notify(event Event) {
	es.mu.RLock()
	defer es.mu.RUnlock()
	for _, subscriber := range es.subscribers {
		subscriber(event)
	}
}

type EventDispatcher struct {
	mu           sync.RWMutex
	eventSources map[string]*EventSource
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		eventSources: make(map[string]*EventSource),
	}
}

func (ed *EventDispatcher) RegisterEventSource(name string) *EventSource {
	ed.mu.Lock()
	defer ed.mu.Unlock()
	eventSource := &EventSource{}
	ed.eventSources[name] = eventSource
	return eventSource
}

func (ed *EventDispatcher) GetEventSource(name string) *EventSource {
	ed.mu.RLock()
	defer ed.mu.RUnlock()
	return ed.eventSources[name]
}

func (ed *EventDispatcher) Dispatch(name string, event Event) {
	eventSource := ed.GetEventSource(name)
	if eventSource != nil {
		eventSource.Notify(event)
	}
}

func main() {
	dispatcher := NewEventDispatcher()

	// 注册一个事件源叫key的事件
	keySource := dispatcher.RegisterEventSource("key")

	// 订阅一个 "create" 事件在key的事件源上
	keySource.Subscribe(func(event Event) {
		if event.Type == "create" {
			date := event.Data.(map[string]interface{})
			fmt.Printf("key created: %v\n", date)
		}
	})

	// 分派一个事件
	event := Event{
		Type: "create",
		Data: map[string]interface{}{
			"key":   "foo",
			"value": "bar",
		},
	}
	dispatcher.Dispatch("key", event)
}
