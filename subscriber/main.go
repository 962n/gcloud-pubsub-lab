package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "pubsub-project-id")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	sub := client.Subscription("MySub")

	err = sub.Receive(ctx, func(ctx context.Context, message *pubsub.Message) {
		fmt.Println(string(message.Data))
		message.Ack()
	})

	if err != nil {
		fmt.Println(err.Error())
	}


}