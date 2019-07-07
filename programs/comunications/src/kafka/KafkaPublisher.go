package kafka

/*
I have seen far by standing on the shoulders of giants.
Example implemented in top of [cluster "github.com/bsm/sarama-cluster"] and ["github.com/segmentio/kafka-go"]
*/
import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"time"
)

/*
In the publisher we create a context with timeout, to make the goroutine that send the event close in 5 seconds
if is not able to finish the process.
*/
func PublishEvents(broker Broker, topic Topic, key string, event string) {
	updatedEvent := event + " and Kafka publisher"
	fmt.Printf("KAFKA Publisher: %s \n", updatedEvent)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	err := createPublisherWriter(broker.Value, topic.Value).WriteMessages(ctx,
		kafka.Message{
			Key:   []byte(key),
			Value: []byte(updatedEvent),
		},
	)
	if err != nil {
		panic(err)
	}
}

func createPublisherWriter(broker string, topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
}
