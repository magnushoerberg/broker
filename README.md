Experimental internal broker for go. Uses RoutingKeys kind of like amqp's topic exchange and queue binding keys.
```
package main

import "fmt"

func main() {
    broker := NewBroker()
    broker.Sub("routing key", func(msg interface{}) {
        fmt.Println(msg)
    })

    broker.Pub("routing key", "hello world")

    broker.Sub("routing key", func(msg interface{}) {
        fmt.Println("second", msg)
    })

    broker.Pub("routing key", []string{"hello world 2", "sdf"})

    broker.Sub("routing", func(msg interface{}) {
        fmt.Println("theird", msg)
    })

    broker.Pub("routing", []string{"hello world 3", "sdf"})
    for {}
}
```

