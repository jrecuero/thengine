package engine_test

import (
	"fmt"
	"testing"

	"github.com/jrecuero/thengine/pkg/engine"
)

func TestMailboxSingleConsumer(t *testing.T) {
	mailbox := engine.GetMailbox()
	if mailbox == nil {
		t.Errorf("[1] GetMailbox Error exp:*Mailbox got:nil")
		return
	}
	mailbox.Clean()

	// create topic
	topicName := "topic/test/1"
	topic := mailbox.CreateTopic(topicName)
	if topic == nil {
		t.Errorf("[1] CreateTopic Error exp:*Topic got:nil")
		return
	}

	// register a consumer to the topic
	consumerName := "consumer/test/1"
	consumer, isNew := mailbox.Subscribe(topicName, consumerName)
	if consumer == nil {
		t.Errorf("[1] Subscribe Error exp:*Consumer got:nil")
		return
	}
	if !isNew {
		t.Errorf("[1] Subscribe Error.isNew exp:%t got:%t", true, isNew)
	}
	ok := mailbox.IsTopicInConsumer(topicName, consumerName)
	if !ok {
		t.Errorf("[2] IsTopicInConsumer Error exp:%t got:%t", true, ok)
	}
	consumer2, isNew2 := mailbox.Subscribe(topicName, consumerName)
	if consumer2 == nil {
		t.Errorf("[1] Re-Subscribe Error exp:*Consumer got:nil")
	}
	if consumer != consumer2 {
		t.Errorf("[1] Re-Subscribe Error.Consumer exp:%v got:%v", consumer, consumer2)
	}
	if isNew2 {
		t.Errorf("[1] Re-Subscribe Error.ResuisNew exp:%t got:%t", true, isNew2)
	}

	// publish one entry into the topic.
	producerName := "producer/test/1"
	messageContent := "topic/sample/1"
	producerMessage := engine.NewMessage(topicName, producerName, consumerName, messageContent)
	errPublish := mailbox.Publish(topicName, producerMessage)
	if errPublish != nil {
		t.Errorf("[1] Publish Error exp:nil got:%+v", errPublish)
	}

	// consume that entry from the topic.
	consumerMessage, err := mailbox.Consume(topicName, consumerName)
	if err != nil {
		t.Errorf("[1] Consume Error exp:nil got:%+v", err)
		return
	}
	if consumerMessage.Topic != topicName {
		t.Errorf("[1] Consumer Message Error.TopicName exp:%s got:%s", topicName, consumerMessage.Topic)
	}
	if consumerMessage.Src != producerName {
		t.Errorf("[1] Consumer Message Error.Source exp:%s got:%s", producerName, consumerMessage.Src)
	}
	if consumerMessage.Dst != consumerName {
		t.Errorf("[1] Consumer Message Error.Destination exp:%s got:%s", consumerName, consumerMessage.Dst)
	}
	message, ok := consumerMessage.Content.(string)
	if !ok {
		t.Errorf("[1] Consumer Message Error exp:string got:error")
		return
	}
	if message != messageContent {
		t.Errorf("[1] Consumer Message Error.Content exp:%s got:%s", messageContent, message)
	}
}

func TestMailboxMultipleConsumer(t *testing.T) {
	mailbox := engine.GetMailbox()
	if mailbox == nil {
		t.Errorf("[2] GetMailbox Error exp:*Mailbox got:nil")
		return
	}
	mailbox.Clean()

	// create topic
	topicName := "topic/test/2"
	topic := mailbox.CreateTopic(topicName)
	if topic == nil {
		t.Errorf("[2] CreateTopic Error exp:*Topic got:nil")
		return
	}

	// register a multiple consumers to the topic
	nbrOfConsumers := 5
	for i := 0; i < nbrOfConsumers; i++ {
		consumerName := fmt.Sprintf("consumer/%d/test/2", i)
		consumer, isNew := mailbox.Subscribe(topicName, consumerName)
		if consumer == nil {
			t.Errorf("[2] Subscribe Error exp:*Consumer got:nil")
			return
		}
		if !isNew {
			t.Errorf("[1] Subscribe Error.isNew exp:%t got:%t", true, isNew)
		}
	}

	// publish one entry into the topic.
	producerName := "producer/test/2"
	messageContent := "topic/sample/2"
	broadcast := "broadcast"
	producerMessage := engine.NewMessage(topicName, producerName, broadcast, messageContent)
	mailbox.Publish(topicName, producerMessage)

	// consume that entry from the topic.
	for i := 0; i < nbrOfConsumers; i++ {
		consumerName := fmt.Sprintf("consumer/%d/test/2", i)
		consumerMessage, err := mailbox.Consume(topicName, consumerName)
		if err != nil {
			t.Errorf("[2] Consume Error exp:nil got:%+v", err)
			return
		}
		if consumerMessage.Topic != topicName {
			t.Errorf("[2] Consumer Message Error.TopicName exp:%s got:%s", topicName, consumerMessage.Topic)
		}
		if consumerMessage.Src != producerName {
			t.Errorf("[2] Consumer Message Error.Source exp:%s got:%s", producerName, consumerMessage.Src)
		}
		if consumerMessage.Dst != broadcast {
			t.Errorf("[2] Consumer Message Error.Destination exp:%s got:%s", broadcast, consumerMessage.Dst)
		}
		message, ok := consumerMessage.Content.(string)
		if !ok {
			t.Errorf("[2] Consumer Message Error exp:string got:error")
			return
		}
		if message != messageContent {
			t.Errorf("[2] Consumer Message Error.Content exp:%s got:%s", messageContent, message)
		}
	}

	// unsubscribe some consumers
	for i := 0; i < nbrOfConsumers/2; i++ {
		consumerName := fmt.Sprintf("consumer/%d/test/2", i)
		err := mailbox.UnSubscribe(topicName, consumerName)
		if err != nil {
			t.Errorf("[2] UnSubscribe Error exp:nil got:%+v", err)
		}
		ok := mailbox.IsTopicInConsumer(topicName, consumerName)
		if ok {
			t.Errorf("[2] IsTopicInConsumer Error.%s exp:%t got:%t", consumerName, false, ok)
		}
	}

	// delete topic
	err := mailbox.DeleteTopic(topicName)
	if err != nil {
		t.Errorf("[2] DeleteTopic Error exp:nil got:%+v", err)
	}
	for i := 0; i < nbrOfConsumers/2; i++ {
		consumerName := fmt.Sprintf("consumer/%d/test/2", i)
		ok := mailbox.IsTopicInConsumer(topicName, consumerName)
		if ok {
			t.Errorf("[2] IsTopicInConsumer Error.%s exp:%t got:%t", consumerName, false, ok)
		}
	}
	topicFound := mailbox.FindTopic(topicName)
	if topicFound != nil {
		t.Errorf("[2] FindTopic Error.%s exp:nil got:%+v", topicName, topicFound)
	}
}

func TestMailboxFaults(t *testing.T) {
	mailbox := engine.GetMailbox()
	if mailbox == nil {
		t.Errorf("[3] GetMailbox Error exp:*Mailbox got:nil")
		return
	}
	mailbox.Clean()

	// re-create mailbox
	mailbox2 := engine.NewMailbox()
	if mailbox2 == nil {
		t.Errorf("[3] NewMailbox Error exp:*Mailbox got:nil")
	}
	if mailbox != mailbox2 {
		t.Errorf("[3] NewMailbox Error exp:%v got:%v", mailbox, mailbox2)
	}

	// subscribe to a not created topic
	topicName := "topic/test/1"
	consumerName := "consumer/test/1"
	consumer, isNew := mailbox.Subscribe(topicName, consumerName)
	if consumer != nil {
		t.Errorf("[3] Subcribe Error exp:nil got:%+v", consumer)
	}
	if isNew {
		t.Errorf("[3] Subscribe Error.isNew exp:%t got:%t", false, isNew)
	}

	// consume from a topic not created
	consumerMessage, err := mailbox.Consume(topicName, consumerName)
	if err == nil {
		t.Errorf("[3] Consume Error exp:error got:nil")
	}
	if consumerMessage != nil {
		t.Errorf("[3] Consume Error.Message exp:nil got:%+v", consumerMessage)
	}

	// delete a topic not created
	err = mailbox.DeleteTopic(topicName)
	if err == nil {
		t.Errorf("[3] DeleteTopic Error exp:error got:nil")
	}

	// publish a message to the wrong topic
	producerName := "producer/test/1"
	messageContent := "topic/sample/1"
	wrongTopic := "topic/test/2"
	newMessage := engine.NewMessage(wrongTopic, producerName, consumerName, messageContent)
	err = mailbox.Publish(topicName, newMessage)
	if err == nil {
		t.Errorf("[3] Publish Error.WrongTopic exp:nil got:%+v", err)
	}
	newMessage = engine.NewMessage(topicName, producerName, consumerName, messageContent)
	err = mailbox.Publish(topicName, newMessage)
	if err == nil {
		t.Errorf("[3] Publish Error.TopicNotFound exp:nil got:%+v", err)
	}

	// unsubscribe to a topic not created
	err = mailbox.UnSubscribe(topicName, consumerName)
	if err == nil {
		t.Errorf("[3] UnSubscribe Error exp:error got:nil")
	}

	// re-create topic
	topic := mailbox.CreateTopic(topicName)
	if topic == nil {
		t.Errorf("[3] CreateTopic Error exp:*Topic got:nil")
		return
	}
	topic2 := mailbox.CreateTopic(topicName)
	if topic2 == nil {
		t.Errorf("[3] Re-CreateTopic Error exp:*Topic got:nil")
	}
	if topic != topic2 {
		t.Errorf("[3] Re-CreateTopic Error.Topic exp:%v got:%v", topic, topic2)
	}

	// remove consumer from topic it was not subcribed
	wrongConsumer := "consumer/test/2"
	err = mailbox.UnSubscribe(topicName, wrongConsumer)
	if err == nil {
		t.Errorf("[3] UnSubscribe Error.WrongConsumer exp:error got:nil")
	}

	// consumer from an empty topic
	consumer, _ = mailbox.Subscribe(topicName, consumerName)
	if consumer == nil {
		t.Errorf("[3] Subscribe Error exp:*Consumer got:nil")
		return
	}
	consumerMessage, err = mailbox.Consume(topicName, consumerName)
	if err != nil {
		t.Errorf("[3] Consume Error exp:nil got:%+v", err)
	}
	if consumerMessage != nil {
		t.Errorf("[3] Consume Error.Empty exp:nil got:%+v", consumerMessage)
	}

	// add consumer to topic again
	consumer, _ = mailbox.Subscribe(topicName, consumerName)
	if consumer == nil {
		t.Errorf("[3] Re-Subscribe Error exp:*Consumer got:nil")
		return
	}

	// consume a message from a consumer not registered
	_, err = mailbox.Consume(topicName, wrongConsumer)
	if err == nil {
		t.Errorf("[3] Consume Error.WrongConsumer exp:error got:nil")
	}
}
