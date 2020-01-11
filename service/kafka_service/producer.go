package kafka_service

import (
	"fmt"
	"gin-server-api/pkg/setting"
	"github.com/Shopify/sarama"
	"log"
	"strings"
	"time"
)

//同步生产者
func SyncProducer() {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Timeout = 5 * time.Second

	p, err := sarama.NewSyncProducer(strings.Split("127.0.0.1:9092", ","), config)
	defer p.Close()

	if err != nil {
		return
	}

	timeData := time.Now()
	value := "this is a message " + timeData.Format("15:04:05")
	defaultTopic := setting.KafkaSetting.DefaultTopic
	msg := &sarama.ProducerMessage{
		Topic: defaultTopic,
	}
	msg.Value = sarama.ByteEncoder(value)

	if _, _, err := p.SendMessage(msg); err != nil {
		log.Fatal(err)
		return
	}
}

//异步生产者
func AsyncProducer() {
	config := sarama.NewConfig()
	//等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//随机向partition发送消息
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	//使用配置 新建一个异步的生产者
	producer, e := sarama.NewAsyncProducer([]string{"127.0.0.1:9092"}, config)
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

	defaultTopic := setting.KafkaSetting.DefaultTopic
	for i := 0; ; i++ {
		time.Sleep(500 * time.Millisecond)
		timeData := time.Now()
		value := "this is a message " + timeData.Format("15:04:05")

		//发动主题消息
		msg := &sarama.ProducerMessage{
			Topic: defaultTopic,
		}

		//字符串转换成字节数组
		msg.Value = sarama.ByteEncoder(value)

		//使用通道发送
		producer.Input() <- msg
	}

}
