package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"google.golang.org/api/iterator"
	"os"
)

func main() {
	ctx := context.Background()

	err := os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:9009")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	client, err := pubsub.NewClient(ctx, "pubsub-project-id")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	topicIterator := client.Topics(ctx)
	mapTopic := map[string]*pubsub.Topic{}
	for {
		topic, iErr := topicIterator.Next()
		if iErr == iterator.Done {
			fmt.Println("done")
			break
		}
		if iErr != nil {
			fmt.Println(iErr.Error())
			return
		}
		mapTopic[topic.ID()] = topic
		fmt.Println(topic.ID())
	}

	topicIDs := []string{
		"MyTopic",
	}

	for _, topicID := range topicIDs {
		_, ok := mapTopic[topicID]
		if ok {
			fmt.Printf("topic:%s already exsit\n", topicID)
			continue
		}
		topic, cErr := client.CreateTopic(ctx, topicID)
		if cErr != nil {
			fmt.Println(err.Error())
			return
		}
		mapTopic[topicID] = topic
	}

	mapSubscription := map[string]*pubsub.Subscription{}
	subIterator := client.Subscriptions(ctx)
	for {
		sub, iErr := subIterator.Next()
		if iErr == iterator.Done {
			fmt.Println("done")
			break
		}
		if iErr != nil {
			fmt.Println(iErr.Error())
			return
		}
		mapSubscription[sub.ID()] = sub
	}

	subscriptions := []struct {
		TopicID string
		SubID   string
	}{
		{TopicID: "MyTopic", SubID: "MySub"},
	}

	for _, sub := range subscriptions {
		_, ok := mapSubscription[sub.SubID]
		if ok {
			fmt.Printf("subscription:%s already exsit\n", sub.SubID)
			continue
		}
		topic, ok := mapTopic[sub.TopicID]
		if !ok {
			fmt.Println("error!!! topic does not exist.")
			return
		}
		config := pubsub.SubscriptionConfig{
			Topic: topic,
		}
		_, cErr := client.CreateSubscription(ctx, sub.SubID, config)
		if cErr != nil {
			fmt.Println(cErr.Error())
		}
	}
}
