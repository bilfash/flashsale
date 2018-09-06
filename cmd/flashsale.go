package main

import (
	"github.com/bilfash/flashsale/interfaces"
	"github.com/bilfash/flashsale/service"
	"github.com/bilfash/thronos"
)

func main() {
	producer := kafka.NewProducer([]string{
		"your kafka broker",
	}, 3, true)
	consumer := kafka.NewConsumer([]string{
		"your kafka broker",
	}, []string{
		"your kafka topic",
	}, "flashsale app", "v1.0")
	cbConfig := thronos.CbConfig{
		ClAddress:   "",
		ClUsername:  "",
		ClPassword:  "",
		BktName:     "",
		BktPassword: "",
	}
	cbHandler := thronos.NewCouchbaseHandler(&cbConfig)
	snapshotRepo := interfaces.NewStockSnapshotRepo(cbHandler)

	service.Consume(consumer, producer, snapshotRepo)
}
