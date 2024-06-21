package ebus

import (
	"reflect"
	"sync"
)

type IEventBus interface {
	Subscribe(subscriber ISubscriber)
	Unsubscribe(subscriber ISubscriber)
	Publish(data interface{}) error
	SubCount() int
	TopicCount() int
}

type EventBus struct {
	subscribers map[string][]ISubscriber
	mutex       sync.Mutex
}

func NewBus() IEventBus {
	return &EventBus{
		subscribers: make(map[string][]ISubscriber),
	}
}

func (eb *EventBus) Subscribe(sub ISubscriber) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	topic := sub.GetTopic()
	if _, ok := eb.subscribers[topic]; !ok {
		eb.subscribers[topic] = make([]ISubscriber, 0)
	}

	eb.subscribers[topic] = append(eb.subscribers[topic], sub)
}

func (eb *EventBus) Unsubscribe(sub ISubscriber) {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()

	topic := sub.GetTopic()
	if subscribers, ok := eb.subscribers[topic]; ok {
		for i := 0; i < len(subscribers); i++ {
			if subscribers[i].ID() == sub.ID() {
				eb.subscribers[topic] = append(subscribers[:i], subscribers[i+1:]...)
				break
			}
		}

		if len(eb.subscribers[topic]) == 0 {
			delete(eb.subscribers, topic)
		}
	}
}

func (eb *EventBus) Publish(data interface{}) error {
	eb.mutex.Lock()
	defer eb.mutex.Unlock()
	topic := reflect.TypeOf(data).String()
	if subscribers, ok := eb.subscribers[topic]; ok {
		for _, subscriber := range subscribers {
			err := subscriber.Consume(data)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (eb *EventBus) TopicCount() int {
	return len(eb.subscribers)
}

func (eb *EventBus) SubCount() int {
	count := 0
	for _, subs := range eb.subscribers {
		count += len(subs)
	}
	return count
}

type EventAssertionError struct {
	Msg string
}

func (e *EventAssertionError) Error() string {
	return e.Msg
}
