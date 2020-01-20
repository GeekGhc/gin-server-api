package kafka_service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
	"sync"
)

var (
	wg sync.WaitGroup
)

func (k *Kafka) Consumer() {
	consumer, err := sarama.NewConsumer(k.host, nil)
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	if err != nil {
		return
	}

	//返回主题的所有分区
	partitionList, err := consumer.Partitions(k.topic)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		pc, err := consumer.ConsumePartition(k.topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			//Message 返回一个消费类型的只读通道 有代理产生
			for msg := range pc.Messages() {
				fmt.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
}
