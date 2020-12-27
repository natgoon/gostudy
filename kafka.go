package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	"github.com/Shopify/sarama"
)

func producer() {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	client, err := sarama.NewClient([]string{"192.168.0.102:9092"}, config)
	defer client.Close()
	if err != nil {
		panic(err)
	}
	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                          sync.WaitGroup
		enqueued, successes, errors int
	)

	wg.Add(1)
	// start a groutines to count successes num
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	// start a groutines to count error num
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			errors++
		}
	}()

ProducerLoop:
	for {
		message := &sarama.ProducerMessage{Topic: "test", Value: sarama.StringEncoder("testing 123")}
		select {
		case producer.Input() <- message:
			enqueued++

		case <-signals:
			producer.AsyncClose() // Trigger a shutdown of the producer.
			break ProducerLoop
		}
	}

	wg.Wait()

	log.Printf("Successfully produced: %d; errors: %d\n", successes, errors)
}

func consumer() {
	fmt.Print("test")
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	client, err := sarama.NewClient([]string{"192.168.0.102:9092"}, config)
	defer client.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("cliend connect success!\n")
	}

	// 创建消费者接口
	consumer, err := sarama.NewConsumerFromClient(client)
	defer consumer.Close()
	if err != nil {
		fmt.Println("consumer: %s\n", err)
	}
	// 获取分区列表
	partitions, err := consumer.Partitions("test")
	if err != nil {
		fmt.Printf("partitions: %s\n", err)
	}

	for _, partitionId := range partitions {
		// 获取各个topic、分区消费者接口
		partitionConsumer, err := consumer.ConsumePartition("test", partitionId, sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("partitionConsumer: %s\n", err)
		}

		go func(pc *sarama.PartitionConsumer) {
			defer (*pc).Close()
			// block
			for message := range (*pc).Messages() {
				value := string(message.Value)
				fmt.Printf("Partitionid: %v; offset:%v,  value: %v\n", message.Partition, message.Offset, value)
			}

		}(&partitionConsumer)

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		select {
		case <-signals:
		}
	}
}

func main() {
	producer()
}
