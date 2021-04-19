package consumer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

const (
	topic         = "testtp1"
	brokerAddress = "localhost:9092"
)

func Consume(ctx context.Context) {

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		GroupID: "my-group",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		fmt.Println("received: ", string(msg.Value))
	}
}