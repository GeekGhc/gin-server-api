package kafka_service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"os"
	"os/signal"
	"sync"
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
	//k.Config.Producer.Partitioner = sarama.NewRandomPartitioner
	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用
	k.Config.Producer.Return.Successes = true
	k.Config.Producer.Return.Errors = true

	//使用配置 新建一个异步的生产者
	producer, err := sarama.NewAsyncProducer(k.host, k.Config)
	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                          sync.WaitGroup
		enqueued, successes, errors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range producer.Successes() {
			fmt.Println(v)
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			errors++
		}
	}()

ProducerLoop:
	for {
		message := &sarama.ProducerMessage{Topic: k.topic, Value: sarama.ByteEncoder(message)}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()
	fmt.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
}
