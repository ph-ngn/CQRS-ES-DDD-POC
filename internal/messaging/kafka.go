package messaging

import (
	"net"
	"strconv"

	"github.com/andyj29/wannabet/internal/log"
	"github.com/segmentio/kafka-go"
)

func NewProducer(address string) *kafka.Writer {
	return &kafka.Writer{
		Addr:        kafka.TCP(address),
		Logger:      kafka.LoggerFunc(log.GetLogger().Infof),
		ErrorLogger: kafka.LoggerFunc(log.GetLogger().Errorf),
		Balancer:    &kafka.Hash{},
	}
}

func NewConsumer(addresses []string, topic, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     addresses,
		Topic:       topic,
		GroupID:     groupID,
		Logger:      kafka.LoggerFunc(log.GetLogger().Infof),
		ErrorLogger: kafka.LoggerFunc(log.GetLogger().Errorf),
	})
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
