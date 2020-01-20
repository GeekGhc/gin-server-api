package kafka_service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"time"
)

//同步生产者
func (k *Kafka) SyncProducer(message string) {
	k.Config.Producer.Return.Successes = true
	k.Config.Producer.Timeout = 5 * time.Second

	producer, err := sarama.NewSyncProducer(k.host, k.Config)

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err != nil {
		panic(err)
	}

	msg := &sarama.ProducerMessage{
		Topic: k.topic,
	}
	//字符串转成字节数组
	msg.Value = sarama.ByteEncoder(message)

	if _, _, err := producer.SendMessage(msg); err != nil {
		log.Fatal(err)
		return
	}
}

//异步生产者
func (k *Kafka) AsyncProducer(message string) {
	//等待服务器所有副本都保存成功后的响应
	k.Config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	k.Config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用
	k.Config.Producer.Return.Successes = true
	k.Config.Producer.Return.Errors = true

	//使用配置 新建一个异步的生产者
	producer, e := sarama.NewAsyncProducer(k.host, k.Config)
	if e != nil {
		return
	}
	defer producer.AsyncClose()

	//判断哪个通道发送的消息
	go func(p sarama.AsyncProducer) {
		for {
			select {
			case suc := <-p.Successes():
				fmt.Println("offset: ", suc.Offset, "timestamp: ", suc.Timestamp.String(), "partitions: ", suc.Partition)
			case fail := <-p.Errors():
				log.Fatal(fail.Err)
			}
		}
	}(producer)

	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		//发动主题消息
		msg := &sarama.ProducerMessage{
			Topic: k.topic,
		}
		//字符串转换成字节数组
		msg.Value = sarama.ByteEncoder(message)
		//使用通道发送
		producer.Input() <- msg
	}
}
