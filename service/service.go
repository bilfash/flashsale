package service

import (
	"encoding/json"
	"fmt"

	kafka2 "github.com/bilfash/kafka-asmara"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"

	"github.com/bilfash/flashsale/interfaces"
	"github.com/bilfash/flashsale/usecase"
)

func Consume(kafkaConsumer kafka2.Consumer, kafkaProducer kafka2.Producer, stockSnapshotRepo *interfaces.StockSnapshotRepo) error {
	consumer := kafkaConsumer.GetConsumer()
	for {
		select {
		case part, ok := <-consumer.Partitions():
			if !ok {
				return fmt.Errorf("%s", "consumer failed to get partition")
			}

			go func(pc.PartitionConsumer) {
				for msg := range pc.Messages() {
					*busy = true
					eventData, err := parseRequest(msg.Value)

					if err != nil {
						fmt.Print("Failed to decode request")
					} else {
						FlashSale(&kafkaProducer, stockSnapshotRepo, *eventData)
						go func() {
							consumer.MarkOffset(msg, "complete")
							consumer.CommitOffsets()
						}()
					}
				}
			}(part)
		}
	}

	return nil
}

func parseRequest(rawData []byte) (message *usecase.StockSnapshot, err error) {
	eventData := &usecase.StockSnapshot{}
	err = json.Unmarshal(rawData, eventData)
	return eventData, err
}
