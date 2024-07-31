package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"time"
)

func main() {
	// to produce messages
	topic := "service_request"
	partition := 0

	kafka.TCP()

	conn, err := kafka.DialLeader(context.Background(), "tcp", kafka.TCP("127.0.0.1:9092").String(), topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.WriteMessages()

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
