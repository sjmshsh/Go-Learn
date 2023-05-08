package main

import (
	"context"
	"fmt"
	"gomock/event"
	"google.golang.org/grpc"
	"log"
	"sync"
)

type EventClient struct {
	mu      sync.Mutex
	sources map[string]*eventSource
	conn    *grpc.ClientConn
}

type Subscriber func(event *event.Event)

type eventSource struct {
	name        string
	stream      event.EventService_SubscribeClient
	subscribers []Subscriber
}

func NewEventClient(address string) (*EventClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := &EventClient{
		sources: make(map[string]*eventSource),
		conn:    conn,
	}

	return client, nil
}

func (c *EventClient) Close() error {
	return c.conn.Close()
}

// Subscribe subscribes to events from the given event source.
func (c *EventClient) Subscribe(source *event.EventSource, subscriber Subscriber) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.sources[source.Name]; ok {
		return fmt.Errorf("event source is already subscribed: %v", source)
	}

	// Create a new event source
	stream, err := event.NewEventServiceClient(c.conn).Subscribe(context.Background(), source)
	if err != nil {
		return err
	}
	eventSource := &eventSource{
		name:   source.Name,
		stream: stream,
	}

	// Add the subscriber to the event source.
	eventSource.subscribers = append(eventSource.subscribers, subscriber)

	// Start a goroutine to receive events from the stream.
	go func() {
		for {
			event, err := stream.Recv()
			if err != nil {
				c.mu.Lock()
				delete(c.sources, source.Name)
				c.mu.Unlock()
				return
			}
			for _, subscriber := range eventSource.subscribers {
				subscriber(event)
			}
		}
	}()

	// Add the event source to the client.
	c.sources[source.Name] = eventSource

	return nil
}

func (c *EventClient) Unsubscribe(source *event.EventSource) error {
	c.mu.Lock()
	c.mu.Unlock()

	// Check if the event source is subscribed.
	eventSource, ok := c.sources[source.Name]
	if !ok {
		return fmt.Errorf("")
	}

	// Close the stream
	if err := eventSource.stream.CloseSend(); err != nil {
		return err
	}
	delete(c.sources, source.Name)

	return nil
}

func main() {
	client, err := NewEventClient(":50051")
	if err != nil {

	}
	defer client.Close()
	// Subscribe to events from the "key" event source.
	source := &event.EventSource{Name: "key"}
	if err := client.Subscribe(source, func(event *event.Event) {
		if event.Type == "create" {
			fmt.Printf("key created: %v\n", event.Data)
		}
	}); err != nil {
		log.Fatalf("failed to subscribe to event source: %v", err)
	}

	// Wait for events.
	select {}
}
