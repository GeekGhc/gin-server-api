package kafka_service

import (
	"gin-server-api/pkg/setting"
	"github.com/Shopify/sarama"
	"strings"
)

var KafkaService *Kafka

type MessageQueue interface {
	SetTopic(topic string)
	SetHost(host string)
}

type Kafka struct {
	host   []string
	topic  string
	Config *sarama.Config
	MessageQueue
}

func SetUp() {
	KafkaService = &Kafka{
		host:   strings.Split(setting.KafkaSetting.Host, ","),
		topic:  setting.KafkaSetting.DefaultTopic,
		Config: sarama.NewConfig(),
	}
}

//host
func (k *Kafka) SetHost(host string) {
	k.host = strings.Split(host, ",")
}

//topic
func (k *Kafka) SetTopic(topic string) {
	if topic != "" {
		k.topic = topic
	} else {
		k.topic = setting.KafkaSetting.DefaultTopic
	}
}

func (k *Kafka) GetTopic() string {
	return k.topic
}
