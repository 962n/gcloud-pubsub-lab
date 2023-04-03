package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"os"
)

func main() {
	err := os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:9009")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	ctx := context.Background()
	client , err := pubsub.NewClient(ctx, "pubsub-project-id")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	topic := client.Topic("MyTopic")
	m := &pubsub.Message{
		Data :[]byte("test"),
	}
	result := topic.Publish(ctx,m)
	id , err := result.Get(ctx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(id)
}
