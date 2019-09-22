package messaging

import (
	"fmt"
	"sync"
	"testing"
)

func sendMessage(b *Broker) {
	go func(b *Broker) {
		message1 := Message{message: "message1"}
		message2 := Message{message: "message2"}
		b.Punish(message1, "one")
		b.Punish(message2, "two")
	}(b)
}

func TestPubSub(t *testing.T) {
	broker := initBroker([]string{"one", "two"})
	defer broker.Close()

	sendMessage(broker)

	var wg sync.WaitGroup

	listen := func(topic string, b *Broker) {
		defer wg.Done()
		message, ok := b.Subscribe(topic)
		if !ok {
			fmt.Printf("[*] Channel closed or not existe\n")
			return
		}
		if message == nil {
			t.Fatal("[*] Get Nil\n")
		}
		fmt.Printf("[*] Get %s from Channel %s\n", message.GetMessage(), topic)
	}

	wg.Add(2)
	go listen("one", broker)
	go listen("two", broker)

	wg.Wait()

}
