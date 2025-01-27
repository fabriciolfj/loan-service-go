package configuration

import (
	"github.com/magiconair/properties"
	"os"
	"strings"
)

type KafkaProperties struct {
	BootstrapServers string
	GroupID          string
	ConsumerTopics   []string
	ProducerTopic    string
	AutoOffsetReset  string
}

func ProvideKafkaProperties() (*KafkaProperties, error) {
	p, err := properties.LoadFile("config.properties", properties.UTF8)
	if err != nil {
		p = properties.NewProperties()
	}

	getValueOrDefault := func(envKey, propKey, defaultValue string) string {
		if envValue := os.Getenv(envKey); envValue != "" {
			return envValue
		}
		return p.GetString(propKey, defaultValue)
	}

	getTopics := func(envKey, configKey, defaultValue string) []string {
		value := getValueOrDefault(envKey, configKey, defaultValue)
		topics := strings.Split(value, ",")
		for i, topic := range topics {
			topics[i] = strings.TrimSpace(topic)
		}
		return topics
	}

	return &KafkaProperties{
		BootstrapServers: getValueOrDefault("KAFKA_BOOTSTRAP_SERVERS", "kafka.bootstrap.servers", "localhost:9092"),
		GroupID:          getValueOrDefault("KAFKA_GROUP_ID", "kafka.group.id", "default-group"),
		ConsumerTopics:   getTopics("KAFKA_CONSUMER_TOPICS", "kafka.topics.consumer", "topic1,topic2"),
		ProducerTopic:    getValueOrDefault("KAFKA_PRODUCER_TOPIC", "kafka.topic.producer", "producer-topic"),
		AutoOffsetReset:  getValueOrDefault("KAFKA_AUTO_OFFSET_RESET", "kafka.auto.offset.reset", "earliest"),
	}, nil
}
