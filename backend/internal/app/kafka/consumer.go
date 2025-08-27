package kafka

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

// Consumer consumes events from Kafka and logs them.
type Consumer struct {
	reader *kafka.Reader
}

// NewConsumer creates a new Kafka consumer for novel.read topic.
func NewConsumer(brokers []string, groupID string) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		GroupID: groupID,
		Topic:   "novel.read",
	})
	return &Consumer{reader: r}
}

// Consume starts consuming messages and logs them.
func (c *Consumer) Consume(ctx context.Context) error {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		log.Printf("novel.read event: %s", string(m.Value))
	}
}

// Close closes the reader.
func (c *Consumer) Close() error {
	return c.reader.Close()
}
