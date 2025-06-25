package kafka

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func SendMessage(topic string, key string, message []byte) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"3.232.44.31:9092"},
		Topic:    topic,
		Balancer: &kafka.Hash{}, // 🔑 clave usada para decidir la partición
	})
	defer writer.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key), // ✅ usar el requestId como clave
		Value: message,
	})
	if err != nil {
		log.Println("Kafka send error:", err)
	}
	return err
}
