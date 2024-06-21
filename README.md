## ebus Example

This example demonstrates how to use the `ebus` package to publish and subscribe to events.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/paulpeters144/ebus"
)

type ExampleEvent struct {
	ID        int
	Content   string
	Author    string
	CreatedAt time.Time
}

func main() {
	// Create a new event bus
	eb := ebus.NewBus()

	// Subscribe to ExampleEvent
	eb.Subscribe(ebus.NewSub(func(data ExampleEvent) {
		fmt.Printf("Received event: %v\n", data)
	}))

	// Publish an event
	event := ExampleEvent{
		ID:        rand.Intn(10_000_000) + 1,
		Content:   "This is example content",
		Author:    "Somebody",
		CreatedAt: time.Now(),
	}

	eb.Publish(event)

	// Unsubscribe from ExampleEvent
	eb.Unsubscribe(sub)

	// Publish another event to demonstrate that unsubscribing works
	eb.Publish(ExampleEvent{
		ID:        rand.Intn(10_000_000) + 1,
		Content:   "This message should not be received",
		Author:    "Somebody Else",
		CreatedAt: time.Now(),
	})

}
```

## License

This project is licensed under the Apache License, Version 2.0. See the [LICENSE](LICENSE) file for details.
