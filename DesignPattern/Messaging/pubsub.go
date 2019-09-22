package messaging

import (
	"log"
	"sync"
)

var (
	once   sync.Once
	broker Broker
)

type (
	// Message struct
	Message struct {
		message string
	}

	// Broker struct
	Broker struct {
		chanMapper map[string]chan Message
	}
)

// GetMessage from Message
func (m *Message) GetMessage() string {
	return m.message
}

// Punish send message to chan
func (b *Broker) Punish(message Message, topic string) {
	if _, ok := b.chanMapper[topic]; ok {
		b.chanMapper[topic] <- message
	}
}

// Subscribe get message from chan
func (b *Broker) Subscribe(topic string) (*Message, bool) {

	if _, ok := b.chanMapper[topic]; !ok {
		return nil, false
	}
	msg, ok := <-b.chanMapper[topic]
	return &msg, ok
}

//UnSubscribe delete from Broker
func (b *Broker) UnSubscribe(topic string) {

	c, ok := b.chanMapper[topic]
	if !ok {
		return
	}
	close(c)
	delete(b.chanMapper, topic)
}

// Close broker
func (b *Broker) Close() {
	for topic, c := range b.chanMapper {
		close(c)
		log.Printf("[*] Close Channel %s.", topic)
	}
}

func initBroker(topicList []string) *Broker {
	once.Do(func() {
		broker = Broker{chanMapper: make(map[string]chan Message)}
		for _, topic := range topicList {
			broker.chanMapper[topic] = make(chan Message)
		}
	})

	return &broker
}
