package main

import (
	"context"
	consumer "kafka-test/consumer"
	producer "kafka-test/producer"
)

func main() {
	ctx := context.Background()

	go producer.Produce(ctx)
	consumer.Consume(ctx)
}