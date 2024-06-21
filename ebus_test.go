package ebus_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/paulpeters144/ebus"
)

type ExampleEvent struct {
	ID        int
	Content   string
	Author    string
	CreatedAt time.Time
}

type ExampleEvent1 struct{}
type ExampleEvent2 struct{}
type ExampleEvent3 struct{}
type ExampleEvent4 struct{}
type ExampleEvent5 struct{}

func TestEventBus(t *testing.T) {
	t.Run("event bus should detect receiving a message", func(t *testing.T) {
		eb := ebus.NewBus()
		received := false
		expectedMessage := "this is example content"
		eb.Subscribe(ebus.NewSub(func(data ExampleEvent) {
			received = data.Content == expectedMessage
		}))

		eb.Publish(ExampleEvent{
			ID:        rand.Intn(10_000_000) + 1,
			Content:   expectedMessage,
			Author:    "Somebody",
			CreatedAt: time.Now(),
		},
		)

		if !received {
			t.Errorf("expected message not received")
		}
	})

	t.Run("event bus should detect receiving a message", func(t *testing.T) {
		eb := ebus.NewBus()
		count := 20
		actualCount := 0
		expectedMessage := "this is example content"
		eb.Subscribe(ebus.NewSub(func(data ExampleEvent) {
			if data.Content == expectedMessage {
				actualCount++
			}
		}))

		for i := 0; i < count; i++ {
			eb.Publish(ExampleEvent{
				ID:        rand.Intn(10_000_000) + 1,
				Content:   expectedMessage,
				Author:    "Somebody",
				CreatedAt: time.Now(),
			},
			)
		}

		if count != actualCount {
			t.Errorf("expected message not received")
		}
	})

	t.Run("event bus should add and remove subscribers", func(t *testing.T) {
		eb := ebus.NewBus()

		sub1 := ebus.NewSub(func(event ExampleEvent) {})
		sub2 := ebus.NewSub(func(event ExampleEvent) {})
		sub3 := ebus.NewSub(func(event ExampleEvent) {})
		sub4 := ebus.NewSub(func(event ExampleEvent) {})
		sub5 := ebus.NewSub(func(event ExampleEvent) {})

		eb.Subscribe(sub1)
		eb.Subscribe(sub2)
		eb.Subscribe(sub3)
		eb.Subscribe(sub4)
		eb.Subscribe(sub5)

		eb.Unsubscribe(sub3)

		len := eb.SubCount()
		if len != 4 {
			t.Errorf("incorrect sub count. should've been 4 but was: %v", len)
		}
	})

	t.Run("should get the correct amout of topics linsted to", func(t *testing.T) {
		eb := ebus.NewBus()

		sub1 := ebus.NewSub(func(event ExampleEvent1) {})
		sub2 := ebus.NewSub(func(event ExampleEvent2) {})
		sub3 := ebus.NewSub(func(event ExampleEvent3) {})
		sub4 := ebus.NewSub(func(event ExampleEvent4) {})

		sub5 := ebus.NewSub(func(event ExampleEvent5) {})
		sub6 := ebus.NewSub(func(event ExampleEvent5) {})
		sub7 := ebus.NewSub(func(event ExampleEvent5) {})
		sub8 := ebus.NewSub(func(event ExampleEvent5) {})

		eb.Subscribe(sub1)
		eb.Subscribe(sub2)
		eb.Subscribe(sub3)
		eb.Subscribe(sub4)
		eb.Subscribe(sub5)

		eb.Subscribe(sub6)
		eb.Subscribe(sub7)
		eb.Subscribe(sub8)

		eb.Unsubscribe(sub3)

		len := eb.TopicCount()
		if len != 4 {
			t.Errorf("incorrect topic count. should've been 4 but was: %v", len)
		}
	})
}
