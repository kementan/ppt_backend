package util

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/IBM/sarama"
	"github.com/gigaflex-co/ppt_backend/config"
)

var producer sarama.SyncProducer
var consumer sarama.Consumer
var WsChannel = make(chan string)
var KfChannel = make(chan string)

func NewKafkaClient(cfg config.Config) {
	brokers := []string{cfg.KafkaBroker}

	configProducer := sarama.NewConfig()
	configProducer.Producer.Return.Successes = true

	var err error
	producer, err = sarama.NewSyncProducer(brokers, configProducer)
	if err != nil {
		log.Fatalf("Error creating producer: %v", err)
	}

	configConsumer := sarama.NewConfig()
	configConsumer.Consumer.Return.Errors = true

	consumer, err = sarama.NewConsumer(brokers, configConsumer)
	if err != nil {
		log.Fatalf("Error creating consumer: %v", err)
	}
}

func ProduceKafkaMessage(topic, message string) {
	kafkaMessage := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	partition, offset, err := producer.SendMessage(kafkaMessage)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
		return
	}

	log.Printf("Message sent to Kafka: %v | %v | %v", kafkaMessage.Topic, partition, offset)
}

func ConsumeKafkaMessage(topic string) {
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
		return
	}
	defer partitionConsumer.Close()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	doneCh := make(chan struct{})

	go func() {
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				message := string(msg.Value)
				KfChannel <- message
			case err := <-partitionConsumer.Errors():
				fmt.Printf("Error: %v\n", err.Err)

			case <-signals:
				close(doneCh)
				return
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
