package ebus

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"reflect"
)

type ISubscriber interface {
	GetTopic() string
	ID() int
	Consume(data interface{}) error
}

type subscriber[T any] struct {
	topic   string
	id      int
	consume func(data T)
}

func NewSub[T any](consume func(data T)) ISubscriber {
	id := rand.Intn(math.MaxInt32) + 1
	sub := &subscriber[T]{
		topic:   reflect.TypeFor[T]().String(),
		id:      id,
		consume: consume}
	return sub
}

func (s *subscriber[T]) GetTopic() string {
	return s.topic
}

func (s *subscriber[T]) ID() int {
	return s.id
}

func (s *subscriber[T]) Consume(data interface{}) error {
	d, err := castEvent[T](data)
	if err != nil {
		return err
	}
	s.consume(d)
	return nil
}

func castEvent[T any](data interface{}) (T, error) {
	castedData, ok := data.(T)
	if !ok {
		t := fmt.Sprintf("Event[%v]", reflect.TypeFor[T]().String())
		d := fmt.Sprintf("%v", data)
		e := fmt.Sprintf("failed to cast %s to type of %s", d, t)
		return castedData, errors.New(e)
	}
	return castedData, nil
}
