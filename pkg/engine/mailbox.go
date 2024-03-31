// mailbox.go contails all structures and logic required to create and handle a
// multiple topics with consumers and producers.
package engine

import (
	"fmt"
	"time"
)

// -----------------------------------------------------------------------------
// Package private variables
// -----------------------------------------------------------------------------

var (
	eMailbox *Mailbox
)

// -----------------------------------------------------------------------------
// Init Package method
// -----------------------------------------------------------------------------

func init() {
	eMailbox = NewMailbox()
}

// -----------------------------------------------------------------------------
// Package public functions
// -----------------------------------------------------------------------------

// GetMailbox function returns the singleton mailbox instance.
func GetMailbox() *Mailbox {
	return eMailbox
}

// -----------------------------------------------------------------------------
//
// Message
//
// -----------------------------------------------------------------------------

// Message structure defines the content for a message sent to the engine
// mailbox.
// Topic: string with the message topic.
// Src: source origen of the message.
// Dst: destination of the message.
// Content: content of the message.
// Time: time when the message was sent.
type Message struct {
	Topic   string
	Src     any
	Dst     any
	Content any
	Time    time.Time
}

// NewMessage function creates a new Message instance with all given
// information.
func NewMessage(topic string, src, dst, content any) *Message {
	time := time.Now()
	return &Message{
		Topic:   topic,
		Src:     src,
		Dst:     dst,
		Content: content,
		Time:    time,
	}
}

// -----------------------------------------------------------------------------
//
// Consumer
//
// -----------------------------------------------------------------------------

// Consumer structure defines the message consumer.
// Name: string with the name of the consumer.
// Pool: list of message the consume is waiting to handle.
type Consumer struct {
	Name string
	pool []*Message
}

// NewConsumer function creates a new Consumer instance with the given name.
func NewConsumer(name string) *Consumer {
	return &Consumer{
		Name: name,
	}
}

// -----------------------------------------------------------------------------
// Consumer Public methods
// -----------------------------------------------------------------------------

// Consume method consumes the first message in the pool of messages.
func (c *Consumer) Consume() *Message {
	if len(c.pool) != 0 {
		message := c.pool[0]
		c.pool = c.pool[1:]
		return message
	}
	return nil
}

// Publish method publishes a new message at the end of the pool of messages.
func (c *Consumer) Publish(message *Message) error {
	c.pool = append(c.pool, message)
	return nil
}

// -----------------------------------------------------------------------------
//
// Topic
//
// -----------------------------------------------------------------------------

// Topic structure defines a new message topic.
// Name: string with the name of the topic.
// Enable: flag to indicate if topic is enable or not.
// Consumers: list of consumers for the topic.
type Topic struct {
	Name      string
	enable    bool
	consumers []*Consumer
}

// NewTopix function creates a new Topic instance with the given name.
func NewTopic(name string) *Topic {
	return &Topic{
		Name:   name,
		enable: true,
	}
}

// -----------------------------------------------------------------------------
// Topic Public methods
// -----------------------------------------------------------------------------

// AddConsumerToTopic method as a consumer to the topic.
// boolean returned shows if the consumer was added to the topic (true) or it
// was already registered (false).
func (t *Topic) AddConsumerToTopic(name string) (*Consumer, bool) {
	if consumer := t.FindConsumer(name); consumer != nil {
		return consumer, false
	}
	consumer := NewConsumer(name)
	t.consumers = append(t.consumers, consumer)
	return consumer, true
}

// Consume method consumes a message from the topic. Every consumer will
// consume that message.
func (t *Topic) Consume(consumerName string) (*Message, error) {
	if consumer := t.FindConsumer(consumerName); consumer != nil {
		return consumer.Consume(), nil
	}
	return nil, fmt.Errorf("Consumer %s not found", consumerName)
}

// FindConsumer method find the Consumer with the given name in the topic.
func (t *Topic) FindConsumer(consumerName string) *Consumer {
	for _, consumer := range t.consumers {
		if consumer.Name == consumerName {
			return consumer
		}
	}
	return nil
}

// Publish method publishes a new message in the topic. Every consumer will
// publish that message.
func (t *Topic) Publish(message *Message) error {
	for _, consumer := range t.consumers {
		if err := consumer.Publish(message); err != nil {
			return err
		}
	}
	return nil
}

// RemoveConsumerFromTopic method removes a consumer from the topic.
func (t *Topic) RemoveConsumerFromTopic(consumerName string) error {
	for i, consumer := range t.consumers {
		if consumer.Name == consumerName {
			t.consumers = append(t.consumers[:i], t.consumers[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("consumer %s not found in topic %s", consumerName, t.Name)
}

// -----------------------------------------------------------------------------
//
// Mailbox
//
// -----------------------------------------------------------------------------

// Mailbox structure defines an engine mailbox.
type Mailbox struct {
	topics    map[string]*Topic
	consumers map[string][]string
}

// NewMailbox function creates a new Mailbox instance.
func NewMailbox() *Mailbox {
	if eMailbox == nil {
		return &Mailbox{
			topics:    make(map[string]*Topic),
			consumers: make(map[string][]string),
		}
	}
	return eMailbox
}

// -----------------------------------------------------------------------------
// Mailbox private methods
// -----------------------------------------------------------------------------

// deleteTopicFromConsumer method deletes the given topic from the list of
// topics of the given consumer.
func (m *Mailbox) deleteTopicFromConsumer(topicName string, consumerName string) {
	if topics, ok := m.consumers[consumerName]; ok {
		for i, tname := range topics {
			if tname == topicName {
				m.consumers[consumerName] = append(m.consumers[consumerName][:i], m.consumers[consumerName][i+1:]...)
				return
			}
		}
	}
}

// -----------------------------------------------------------------------------
// Mailbox Public methods
// -----------------------------------------------------------------------------

// Clean method cleans the mailbox with brand new and emtpy topics and
// consumers.
func (m *Mailbox) Clean() {
	m.topics = make(map[string]*Topic)
	m.consumers = make(map[string][]string)
}

// Consume method consumes a message for the given topic and the given consumer.
func (m *Mailbox) Consume(topicName string, consumerName string) (*Message, error) {
	if topic := m.FindTopic(topicName); topic != nil {
		return topic.Consume(consumerName)
	}
	return nil, fmt.Errorf("Consumer %s not found", consumerName)
}

// CreateTopic method creates a new topic.
func (m *Mailbox) CreateTopic(name string) *Topic {
	if topic := m.FindTopic(name); topic != nil {
		return topic
	}
	topic := NewTopic(name)
	m.topics[name] = topic
	return topic
}

// DeleteTopic method deletes a topic.
func (m *Mailbox) DeleteTopic(topicName string) error {
	if topic := m.topics[topicName]; topic != nil {
		var consumers []string
		for _, consumer := range topic.consumers {
			consumers = append(consumers, consumer.Name)
		}
		delete(m.topics, topicName)
		for _, consumerName := range consumers {
			m.deleteTopicFromConsumer(topicName, consumerName)
		}
		return nil
	}
	return fmt.Errorf("topic %s not found", topicName)
}

// FindTopic method finds a topic with the given name.
func (m *Mailbox) FindTopic(name string) *Topic {
	return m.topics[name]
}

// IsTopicInConsumer method finds the given consumer in the list of consumers
// for the given topic.
func (m *Mailbox) IsTopicInConsumer(topicName string, consumerName string) bool {
	if topics, ok := m.consumers[consumerName]; ok {
		for _, tname := range topics {
			if tname == topicName {
				return true
			}
		}
	}
	return false
}

// Publish method publishes a message for a give topic.
func (m *Mailbox) Publish(topicName string, message *Message) error {
	if message.Topic != topicName {
		return fmt.Errorf("topic %s does not match message topic %s", topicName, message.Topic)
	}
	if topic := m.FindTopic(topicName); topic != nil {
		return topic.Publish(message)
	}
	return fmt.Errorf("topic %s not found", topicName)
}

// Subscribe method subscribe a consumer to a topic.
func (m *Mailbox) Subscribe(topicName string, consumerName string) (*Consumer, bool) {
	if topic := m.FindTopic(topicName); topic != nil {
		if consumer := topic.FindConsumer(consumerName); consumer != nil {
			return consumer, false
		}
		consumer, isNew := topic.AddConsumerToTopic(consumerName)
		if isNew {
			m.consumers[consumerName] = append(m.consumers[consumerName], topicName)
		}
		return consumer, isNew
	}
	return nil, false
}

// UnSubscribe method unsubscribe a consumer from a topic.
func (m *Mailbox) UnSubscribe(topicName string, consumerName string) error {
	if topic := m.FindTopic(topicName); topic != nil {
		m.deleteTopicFromConsumer(topicName, consumerName)
		return topic.RemoveConsumerFromTopic(consumerName)
	}
	return fmt.Errorf("topic %s not found", topicName)
}
