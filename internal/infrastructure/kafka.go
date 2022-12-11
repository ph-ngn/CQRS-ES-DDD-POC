package infrastructure

import (
	"context"
	"encoding/json"
	"net"
	"strconv"

	"github.com/andyj29/wannabet/internal/domain/common"
	kafka "github.com/segmentio/kafka-go"
)

func NewProducer(address, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:        kafka.TCP(address),
		Topic:       topic,
		Logger:      kafka.LoggerFunc(infraLogger.Infof),
		ErrorLogger: kafka.LoggerFunc(infraLogger.Errorf),
		Balancer:    &kafka.Hash{},
	}
}

func NewConsumer(addresses []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     addresses,
		Topic:       topic,
		GroupID:     groupID,
		Logger:      kafka.LoggerFunc(infraLogger.Infof),
		ErrorLogger: kafka.LoggerFunc(infraLogger.Errorf),
	})
}

type eventBus struct {
	producer *kafka.Writer
}

func NewAsyncEventBus(producer *kafka.Writer) *eventBus {
	return &eventBus{
		producer: producer,
	}
}

func (b *eventBus) Publish(event common.Event) {
	payload, err := json.Marshal(event)
	if err != nil {
		return
	}

	if err := b.producer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(event.GetAggregateID()),
			Value: payload,
		},
	); err != nil {
		// To implement storage fallback and retries
		return
	}
}

type topicConfig struct {
	topic         string
	numPartitions int
	repFactor     int
}

func NewTopicConfig(topic string, numPartitions, repFactor int) *topicConfig {
	return &topicConfig{
		topic:         topic,
		numPartitions: numPartitions,
		repFactor:     repFactor,
	}
}

func CreateTopics(protocol, address string, topicConfigs ...topicConfig) error {
	conn, err := kafka.Dial(protocol, address)
	if err != nil {
		return err
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		return err
	}

	controllerConn, err := kafka.Dial(protocol, net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		return err
	}
	defer controllerConn.Close()

	kafkaTopicConfigs := make([]kafka.TopicConfig, 0)
	for _, config := range topicConfigs {
		cfg := kafka.TopicConfig{
			Topic:             config.topic,
			NumPartitions:     config.numPartitions,
			ReplicationFactor: config.repFactor,
		}
		kafkaTopicConfigs = append(kafkaTopicConfigs, cfg)
	}
	return controllerConn.CreateTopics(kafkaTopicConfigs...)
}
