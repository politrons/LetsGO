package connectors

/*
I have seen far by standing on the shoulders of giants.
Example implemented in top of [cluster "github.com/bsm/sarama-cluster"] and ["github.com/segmentio/kafka-go"]
*/

import (
	"context"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
	"testing"
	"time"
)

/*
To create a Kafka consumer it require three different blocks.
* Configuration: Using [cluster.NewConfig()] we are able to create a simple config with all config by default
* Consumer: Using [cluster.NewConsumer] we are able to create a Consumer where we need to specify, array of brokers,
groupId, array of topics, and finally the config created previously.
* Listener: Using [consumer.Messages()] inside a range, it create a channel of [ConsumerMessage].
It block the channel per iteration of the range, until a message arrive from the channel.

Then we receive [ConsumerMessage], which contains headers, key/value, partition info, offset.
Using the switch operator with consumerMessage.Topic we route the message into the specific topic handler
and there we process the message
*/
const topic = "MyTopic"
const broker = "localhost:9092"

func TestKafkaConsumer(t *testing.T) {
	config := cluster.NewConfig()
	consumer, err := createConsumer(config)
	if err != nil {
		panic(err)
	}
	consumerListener(consumer)
}

func createConsumer(config *cluster.Config) (*cluster.Consumer, error) {
	return cluster.NewConsumer(
		[]string{broker},
		"group-id",
		[]string{topic},
		config)
}

func consumerListener(consumer *cluster.Consumer) {
	for consumerMessage := range consumer.Messages() {
		switch consumerMessage.Topic {
		case topic:
			println(string(consumerMessage.Value))
			break
		default:
			println("Error kafka message not expected for any Topic")
		}
	}
}

/*
To create a Kafka producer it require two different blocks.
* Configuration: Using [kafka.NewWriter] we create a struct that need to contain, array of brokers, topic abd
strategy of how to distribute the message. [LeastBytes] is a Balancer implementation that routes messages to the partition
that has received the least amount of data.

* Producer: Using [publisher.WriteMessages] we pass a context to specify timeout of the process to send the message,
and then, the number of message that we want to send using [kafka.Message] struct where you can specify header, key/value,
partition, offset.
*/
func TestKafkaProducer(t *testing.T) {
	publisher := createPublisherWriter()
	defer publisher.Close()
	publishEvents(publisher)
}

func createPublisherWriter() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}

/*
In the publisher we create a context with timeout, to make the goroutine that send the event close in 5 seconds
if is not able to finish the process
*/
func publishEvents(publisher *kafka.Writer) {
	for {
		uuidString, _ := uuid.NewRandom()
		ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
		err := publisher.WriteMessages(ctx,
			kafka.Message{
				Key:   []byte("Key-" + uuidString.String()),
				Value: []byte(uuidString.String()),
			},
		)
		if err != nil {
			panic(err)
		}
	}
}
