package asyncbus

import (
	"context"
	"encoding/json"
	"net"
	"strconv"

	"github.com/andyj29/wannabet/internal/domain/common"
	"github.com/andyj29/wannabet/internal/infrastructure/logger"
	kafka "github.com/segmentio/kafka-go"
)

func NewProducer(address, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:        kafka.TCP(address),
		Topic:       topic,
		Logger:      kafka.LoggerFunc(logger.InfraLogger.Infof),
		ErrorLogger: kafka.LoggerFunc(logger.InfraLogger.Errorf),
		Balancer:    &kafka.Hash{},
	}
}

func NewConsumer(addresses []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     addresses,
		Topic:       topic,
		GroupID:     groupID,
		Logger:      kafka.LoggerFunc(logger.InfraLogger.Infof),
		ErrorLogger: kafka.LoggerFunc(logger.InfraLogger.Errorf),
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

type TopicConfig struct {
	topic         string
	numPartitions int
	repFactor     int
}

func CreateTopics(protocol, address string, topicConfigs ...TopicConfig) error {
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
