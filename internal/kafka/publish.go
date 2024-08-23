package kafka

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
	log "github.com/sirupsen/logrus"
	"os"
)

type Publisher struct {
	logger      *log.Logger
	kafkaWriter kafka.Writer
}

func NewPublisher(l *log.Logger, topic string) Publisher {
	mechanism, _ := scram.Mechanism(scram.SHA256, os.Getenv("KAFKA_USERNAME"), os.Getenv("KAFKA_PASSWORD"))

	return Publisher{
		logger: l,
		kafkaWriter: kafka.Writer{
			Addr:  kafka.TCP(os.Getenv("KAFKA_HOST")),
			Topic: topic,
			Transport: &kafka.Transport{
				SASL: mechanism,
				TLS:  &tls.Config{},
			},
		},
	}
}

func (p *Publisher) Publish(data interface{}) {
	payload, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling games to JSON: %v\n", err)
	}

	p.logger.WithFields(log.Fields{
		"payload":    string(payload[:30]),
		"sourceFile": "publish.go",
		"function":   "publish",
	}).Info("publishing message")

	err = p.kafkaWriter.WriteMessages(context.Background(), kafka.Message{Value: payload})
	if err != nil {
		fmt.Println(err)
	}

	err = p.kafkaWriter.Close()
	if err != nil {
		fmt.Println(err)
	}
}
