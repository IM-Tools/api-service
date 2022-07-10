/**
  @author:panliang
  @data:2022/6/6
  @note
**/
package kafak

import (
	"fmt"
	"gopkg.in/Shopify/sarama.v1"
	dao2 "im-services/app/api/service/dao"
	"im-services/config"
	"im-services/service/dao"
	"sync"
)

var (
	offlineMessageDao *dao2.OfflineMessageDao
)

func ConsumerInit() {

	offlineMessageDao = dao.New()

	var wg sync.WaitGroup
	consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s:%s", config.Conf.Kafka.Host, config.Conf.Kafka.Port)}, nil)
	if err != nil {
		fmt.Println("Failed to start consumer: %s", err)
		return
	}
	partitionList, err := consumer.Partitions("web_log") //获得该topic所有的分区
	if err != nil {
		fmt.Println("Failed to get the list of partition:, ", err)
		return
	}
	fmt.Println(partitionList)

	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("web_log", int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Println("Failed to start consumer for partition %d: %s\n", partition, err)
			return
		}
		wg.Add(1)
		go func(sarama.PartitionConsumer) { //为每个分区开一个go协程去取值
			for msg := range pc.Messages() { //阻塞直到有值发送过来，然后再继续等待

				offlineMessageDao.PrivateOfflineMessageSave(string(msg.Value))

				fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
			defer pc.AsyncClose()
			wg.Done()
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}
