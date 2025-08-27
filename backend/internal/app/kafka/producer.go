package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

// Producer publishes events to Kafka.
type Producer struct {
	writer *kafka.Writer
}

// NewProducer creates a new Kafka producer.
func NewProducer(brokers []string) *Producer {
	return &Producer{writer: &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    "novel.read",
		Balancer: &kafka.LeastBytes{},
	}}
}

// ReadEvent represents a chapter read event.
type ReadEvent struct {
	NovelID   uint `json:"novel_id"`
	ChapterID uint `json:"chapter_id"`
}

// PublishNovelRead sends a novel.read event to Kafka.
func (p *Producer) PublishNovelRead(ctx context.Context, novelID, chapterID uint) error {
	evt := ReadEvent{NovelID: novelID, ChapterID: chapterID}
	b, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	return p.writer.WriteMessages(ctx, kafka.Message{Value: b})
}

// Close shuts down the producer.
func (p *Producer) Close() error {
	return p.writer.Close()
}
