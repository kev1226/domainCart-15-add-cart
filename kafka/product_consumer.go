package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"log"
	"time"

	kafkaLib "github.com/segmentio/kafka-go"
)

// 💡 Calcula la partición a partir del requestID
func partitionForKey(key string, numPartitions int) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32()) % numPartitions
}

func WaitForProductResponse(requestID string) (*ProductResponse, error) {
	const partitionCount = 3
	partition := partitionForKey(requestID, partitionCount)

	log.Printf("🧭 Esperando respuesta para requestID=%s en partición %d", requestID, partition)

	reader := kafkaLib.NewReader(kafkaLib.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       ProductResponseTopic,
		Partition:   partition,
		MinBytes:    1,
		MaxBytes:    10e6,
		StartOffset: kafkaLib.LastOffset,
	})
	defer reader.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	start := time.Now()

	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				log.Printf("⏰ Timeout Kafka (%.2fs)", time.Since(start).Seconds())
				return nil, fmt.Errorf("producto no encontrado o tiempo agotado")
			}
			log.Printf("⚠️ Error leyendo Kafka: %v", err)
			continue
		}

		log.Printf("📥 Mensaje recibido de Kafka en partición %d: %s", partition, string(msg.Value))

		var res ProductResponse
		if err := json.Unmarshal(msg.Value, &res); err != nil {
			log.Printf("⚠️ Error de deserialización: %v", err)
			continue
		}

		log.Printf("🔍 Comparando requestId recibido %s con esperado %s", res.RequestID, requestID)

		if res.RequestID == requestID && res.ID != 0 {
			log.Printf("✅ Producto recibido en %.2fs en partición %d", time.Since(start).Seconds(), partition)
			return &res, nil
		}

		log.Printf("⚠️ Mensaje ignorado: ID=%v, RequestID=%s", res.ID, res.RequestID)
	}
}
